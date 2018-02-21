package config

import (
	"fmt"
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

type Configuration struct {
	Storage
	Debug                 bool
	Port                  int    `default:"8787" validate:"min=1024,max=65535"`
	ServiceProviderConfig string `validate:"omitempty,pathexists,isfile=.json"`
	Config                string `validate:"omitempty,pathexists,isdir"` // (todo) > check the config directory contains two directories, one for resource types and one for schemas, and that them contains json files
	PageSize              int    `default:"10" validate:"min=1,max=10"`
	Enable
}

type Storage struct {
	Type string `default:"mongo" validate:"eq=mongo"` // (note) > since we are only supporting mongo at the moment
	Host string `default:"0.0.0.0" validate:"hostname|ip4_addr"`
	Port int    `default:"27017" validate:"min=1024,max=65535"`
	Name string `default:"scimd" validate:"min=1,excludesall=/\\.*<>:?$\""` // cannot contain any of these characters /, \, ., *, <, >, :, , ?, $, " (fixme) exclude also => |
	Coll string `default:"resources" validate:"min=1,excludes=$,nstartswith=system."`
}

// (fixme) > can make this pvt?

type Enable struct {
	Self bool
}

// (todo) > make getters for the following variables (making them pvt to config)

// Values contains the configuration values
var Values *Configuration

// Errors contains the happened configuration errors
var Errors validator.ValidationErrors

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
		errors, _ := err.(validator.ValidationErrors)
		Errors = append(Errors, errors...)
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

	// (todo) > check wheter custom configs are given, in such case override the defaults one
	// 2. clean, then new push on repositories
	// 1. override svc variable
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
