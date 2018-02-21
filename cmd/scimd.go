package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/defaults"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/cobra"
)

var scimd = &cobra.Command{
	Use:   "scimd",
	Short: "SCIMD is ...",
	Long: `Long description here ...
	
Bla bla ...
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
	},
	Run: func(cmd *cobra.Command, args []string) {
		spew.Dump(defaults.ServiceProviderConfig)
		fmt.Println()
		spew.Dump(core.GetSchemaRepository().List())
		fmt.Println()
		spew.Dump(core.GetResourceTypeRepository().List())
		fmt.Println()
		spew.Dump(config.Values)

		server.Get(defaults.ServiceProviderConfig).Run(":" + strconv.Itoa(config.Values.Port))
	},
}

func init() {
	// Definition of flags
	scimd.Flags().IntVarP(&config.Values.Port, "port", "p", config.Values.Port, "port to run the server on")
	scimd.PersistentFlags().BoolVar(&config.Values.Debug, "debug", config.Values.Debug, "wheter to enable or not the debug mode")

	// Binding flags to configuration manager
	viper.BindPFlags(scimd.Flags())
}

// Execute is ...
func Execute() {
	scimd.SuggestionsMinimumDistance = 1
	if err := scimd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
