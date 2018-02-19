package config

import (
	"strings"

	"github.com/fabbricadigitale/scimd/defaults"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fatih/structs"
	d "github.com/mcuadros/go-defaults"

	"github.com/spf13/viper"
)

type Configuration struct {
	Storage
	ServiceProviderConfig string
	Config                string
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
	d.SetDefaults(Values)

	vip := viper.New()
	for key, value := range structs.Map(Values) {
		vip.SetDefault(key, value)
	}
	vip.SetConfigName(filename)
	vip.AddConfigPath(".")

	vip.SetEnvPrefix("scimd")
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.AutomaticEnv()

	var err error
	err = vip.ReadInConfig()
	err = vip.Unmarshal(&Values)

	return err
}

func init() {
	getConfig(".scimd")

	// Schemas
	core.GetSchemaRepository().Push(defaults.UserSchema)
	core.GetSchemaRepository().Push(defaults.GroupSchema)

	// Resource types
	core.GetResourceTypeRepository().Push(defaults.UserResourceType)
	core.GetResourceTypeRepository().Push(defaults.GroupResourceType)
}

// (todo)
// OVERRIDE ALL CONFIG
// scimd --service-provider-config <path> --config <path>

// GET ALL CONFIG
// Via static command
// scimd get-config <path> => download config directory <within path>
// scimd get-service-provider-config <path> => download service provider config file within <path>
