package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fabbricadigitale/scimd/validation"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/fabbricadigitale/scimd/defaults"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fatih/structs"
	d "github.com/mcuadros/go-defaults"
)

// Configuration is ...
type Configuration struct {
	Storage
	Debug                 bool
	Port                  int    `default:"8787" validate:"min=1024,max=65535"`
	ServiceProviderConfig string `validate:"omitempty,pathexists,isfile=.json"`
	Config                string `validate:"omitempty,pathexists,isdir,hassubdir=schemas,hassubdir=resources"` // (todo) > check the config directory contains two directories, one for resource types and one for schemas, and that them contains json files
	PageSize              int    `default:"10" validate:"min=1,max=10"`
	Enable
}

// Storage is ...
type Storage struct {
	Type string `default:"mongo" validate:"eq=mongo"` // (note) > since we are only supporting mongo at the moment
	Host string `default:"0.0.0.0" validate:"hostname|ip4_addr"`
	Port int    `default:"27017" validate:"min=1024,max=65535"`
	Name string `default:"scimd" validate:"min=1,excludesall=/\\.*<>:?$\""` // cannot contain any of these characters /, \, ., *, <, >, :, , ?, $, " (fixme) exclude also => |
	Coll string `default:"resources" validate:"min=1,excludes=$,nstartswith=system."`
}

// Enable is ...
type Enable struct {
	Self bool
}

var (
	// Values contains the configuration values
	Values *Configuration
	// Errors contains the happened configuration errors
	Errors validator.ValidationErrors
)

var serviceProviderConfig core.ServiceProviderConfig

// getConfig is responsible to set configuration values
//
// The priority model from higher to lower is the following one.
// 0. Flags
// 1. Environment variables
// 2. Configuration file
func getConfig(filename string) {
	Values = new(Configuration)

	// Defaults
	d.SetDefaults(Values)
	for key, value := range structs.Map(Values) {
		viper.SetDefault(key, value)
	}

	viper.SetConfigName(filename)

	viper.AddConfigPath(".")
	// Search home directory
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(home)

	viper.SetEnvPrefix("scimd")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	err = viper.Unmarshal(&Values)
	// (todo) > better handling of errors - invalid syntax - etc. etc.
	// exit(1) ?
	if err != nil {
		fmt.Println("xxxxxx")
		panic(err)
	}

	// Validate the configurations and collect errors
	_, err = Valid()
	if err != nil {
		errs, _ := err.(validator.ValidationErrors)
		Errors = append(Errors, errs...)
	}
}

func init() {
	getConfig(".scimd")

	// ServiceProviderConfig
	serviceProviderConfig = defaults.ServiceProviderConfig

	// Schemas
	core.GetSchemaRepository().Push(defaults.UserSchema)
	core.GetSchemaRepository().Push(defaults.GroupSchema)

	// Resource types
	core.GetResourceTypeRepository().Push(defaults.UserResourceType)
	core.GetResourceTypeRepository().Push(defaults.GroupResourceType)
}

func getFilesWithExt(dir string, ext string) (files []string) {
	re := regexp.MustCompile(`^\.?(.*)`)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if filepath.Ext(path) == re.ReplaceAllString(ext, `.$1`) {
				files = append(files, path)
			}
		}
		return nil
	})
	return
}

// Custom is responsible to load custom configurations
//
// First it checks whether user have specified custom configs.
// Then checks whether them are suitable and valid.
// Finally tries to load them.
func Custom() error {
	// Check wheter user specified a custom service provider config
	if Values.ServiceProviderConfig != "" {
		dat, _ := ioutil.ReadFile(Values.ServiceProviderConfig)
		serviceProviderConfig = *core.NewServiceProviderConfig()
		if err := json.Unmarshal(dat, &serviceProviderConfig); err != nil {
			if Values.Debug {
				fmt.Fprintf(os.Stderr, err.Error())
			}
			return fmt.Errorf("Error unmarshalling custom service provider config (\"%s\")", Values.ServiceProviderConfig)
		}
		// Here default service provider config has been overridden

		if errs := validation.Validator.Struct(serviceProviderConfig); errs != nil {
			return fmt.Errorf(validation.Errors(errs))
		}
	}

	// Check wheter user specified a custom location to provide its own resources (schemas + resource types)
	if Values.Config != "" {
		schemas := getFilesWithExt(filepath.Join(Values.Config, "schemas"), "json")
		rstypes := getFilesWithExt(filepath.Join(Values.Config, "resources"), "json")

		// (todo) > check correspondance between schemas and resource types or this is user's responsibility?

		core.GetSchemaRepository().Clean()
		for _, schema := range schemas {
			_, err := core.GetSchemaRepository().PushFromFile(schema)
			if err != nil {
				return fmt.Errorf("Error loading schema (\"%s\")", schema)
			}
		}

		core.GetResourceTypeRepository().Clean()
		for _, rstype := range rstypes {
			_, err := core.GetResourceTypeRepository().PushFromFile(rstype)
			if err != nil {
				return fmt.Errorf("Error loading resource type (\"%s\")", rstype)
			}
		}
	}

	return nil
}

// Valid checks wheter the configuration is valid or not
func Valid() (bool, error) {
	if err := validation.Validator.Struct(Values); err != nil {
		return false, err
	}
	return true, nil
}

// ServiceProviderConfig returns the current service provider config
func ServiceProviderConfig() core.ServiceProviderConfig {
	return serviceProviderConfig
}
