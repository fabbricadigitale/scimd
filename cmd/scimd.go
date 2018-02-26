package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/cobra"
)

var scimd = &cobra.Command{
	Use: "scimd",
	// TraverseChildren: true,
	Short: "SCIMD is ...",
	Long: `...
	
	SCIM 2 RFC - published under the IETF - defines an open API for managing identities.
Complete documentation is available at ...`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Flags overrides (or merge with) configuration values
		// Thus we re-validate the configuration prior to execute and we collect errors
		if _, err := config.Valid(); err != nil {
			errors, _ := err.(validator.ValidationErrors)
			config.Errors = append(config.Errors, errors...)
		}

		// Printing errors
		if len(config.Errors) > 0 {
			fmt.Fprintln(os.Stderr, validation.Errors(config.Errors))
			os.Exit(1)
		}

		// Ensure eventual custom config is ok and load it
		if err := config.Custom(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Start the server with the current service provider config
		spc := config.ServiceProviderConfig()
		server.Get(&spc).Run(":" + strconv.Itoa(config.Values.Port))
	},
	DisableAutoGenTag: true,
}

func init() {
	// Definition of flags
	scimd.Flags().StringVar(&config.Values.Storage.Type, "storage-type", config.Values.Storage.Type, "type of storage to use")
	scimd.Flags().StringVar(&config.Values.Storage.Host, "storage-host", config.Values.Storage.Host, "the storage's address")
	scimd.Flags().IntVar(&config.Values.Storage.Port, "storage-port", config.Values.Storage.Port, "the storage's port")
	scimd.Flags().StringVar(&config.Values.Storage.Name, "storage-name", config.Values.Storage.Name, "the storage's database name")
	scimd.Flags().StringVar(&config.Values.Storage.Coll, "storage-coll", config.Values.Storage.Coll, "the storage's collection name")

	scimd.PersistentFlags().BoolVar(&config.Values.Debug, "debug", config.Values.Debug, "wheter to enable or not the debug mode")

	scimd.Flags().IntVarP(&config.Values.Port, "port", "p", config.Values.Port, "port to run the server on")

	scimd.Flags().StringVarP(&config.Values.ServiceProviderConfig, "service-provider-config", "s", config.Values.ServiceProviderConfig, "the path of the service provider config to use")
	scimd.Flags().StringVarP(&config.Values.Config, "config", "c", config.Values.Config, "the path of directory containing the configuration resources")

	scimd.Flags().IntVar(&config.Values.PageSize, "page-size", config.Values.PageSize, "the page size the server has to use")

	scimd.Flags().BoolVar(&config.Values.Enable.Self, "enable-self", config.Values.Enable.Self, "whether to enable or not the creation of an endpoint for the authenticated entity")

	// Binding flags to configuration manager
	viper.BindPFlags(scimd.Flags())
}

// Execute is starting point for commands
func Execute() {
	scimd.SuggestionsMinimumDistance = 1
	if err := scimd.Execute(); err != nil {
		os.Exit(1)
	}
}
