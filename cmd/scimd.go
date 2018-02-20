package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/defaults"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/spf13/cobra"
)

var scimdCmd = &cobra.Command{
	Use:   "scimd",
	Short: "SCIMD is ...",
	Long: `Long description here ...
	
Bla bla ...
Complete documentation is available at ...`,
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

// Execute is ...
func Execute() {
	scimdCmd.SuggestionsMinimumDistance = 1
	if err := scimdCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
