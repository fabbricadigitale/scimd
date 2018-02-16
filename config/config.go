package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	"gopkg.in/go-playground/validator.v9"

	"github.com/fatih/structs"
	"github.com/mcuadros/go-defaults"

	"github.com/fabbricadigitale/scimd/schemas/core"
	v "github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/viper"
)

type Configuration struct {
	Storage
	Resources
}

type Resources struct {
	Dir string `default:"default"`
}

type Storage struct {
	Type string `default:"mongo"`
	Host string `default:"0.0.0.0"`
	Port string `default:"27017"`
	Name string `default:"scimd"`
	Coll string `default:"resources"`
}

var Values *Configuration

func getConfig(filename string) error {
	Values = new(Configuration)
	defaults.SetDefaults(Values)

	vip := viper.New()
	for key, value := range structs.Map(Values) {
		vip.SetDefault(key, value)
	}
	vip.SetConfigName(filename)
	vip.AddConfigPath(".")

	vip.SetEnvPrefix(".scimd")
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.AutomaticEnv()

	var err error
	err = vip.ReadInConfig()
	err = vip.Unmarshal(&Values)

	return err
}

func getAssets(dir string, fx func(data []byte)) {
	box := packr.NewBox(filepath.Join(Values.Resources.Dir, dir))
	list := box.List()

	for _, p := range list {
		bytes, err := box.MustBytes(p)
		if err != nil {
			panic(err) // (fixme) > panic or exit?
		}
		fx(bytes)
	}
}

func init() {
	getConfig("scimd")
}

// Get returns the service provider configuration.
//
// It also populates the repositories.
func Get() *core.ServiceProviderConfig {
	// Need this to enforce packr to generate go codes for the default location
	packr.NewBox("default")
	packr.NewBox("default/schemas")
	packr.NewBox("default/resources")

	// Service provider configuration
	box := packr.NewBox(Values.Resources.Dir)
	dat, err := box.MustBytes("service_provider_config.json")
	if err != nil {
		panic(err) // (fixme) > panic or exit?
	}
	spc := core.NewServiceProviderConfig()
	err = json.Unmarshal(dat, &spc)
	if err != nil {
		panic("unmarshalling errors")
	}

	if err := v.Validator.Struct(spc); err != nil {
		errs := ""
		for _, e := range err.(validator.ValidationErrors) {
			errs += fmt.Sprintf("%s\n", e)
		}
		panic(fmt.Sprintf("validation errors\n%s", errs)) // (fixme) > panic or exit?
	}

	// Schemas
	getAssets("schemas", func(data []byte) {
		core.GetSchemaRepository().PushFromData(data)
	})

	// Resource types
	getAssets("resources", func(data []byte) {
		core.GetResourceTypeRepository().PushFromData(data)
	})

	return spc
}
