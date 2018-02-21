package cmd

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/spf13/cobra"
)

func init() {
	scimd.AddCommand(printConfig)
}

var printConfig = &cobra.Command{
	Use:   "print-config",
	Short: "Print the current configuration",
	Long:  `Explicitly dump the current configuration objects for debuggin purposes`,
	Run: func(cmd *cobra.Command, args []string) {
		dump := spew.NewDefaultConfig()
		dump.DisablePointerAddresses = true
		dump.SortKeys = true

		fmt.Fprintln(os.Stdout, "SERVICE PROVIDER CONFIG")
		dump.Dump(config.ServiceProviderConfig())

		fmt.Fprintln(os.Stdout, "\nSCHEMAS")
		dump.Dump(core.GetSchemaRepository().List())

		fmt.Fprintln(os.Stdout, "\nRESOURCE TYPES")
		dump.Dump(core.GetResourceTypeRepository().List())

		fmt.Fprintln(os.Stdout, "\nCONFIGURATION VALUES")
		dump.Dump(config.Values)
	},
}
