package cmd

import (
	"fmt"
	"os"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/cobra"
)

func init() {
	scimd.AddCommand(getServiceProviderConfigCmd)
}

var getServiceProviderConfigCmd = &cobra.Command{
	Use:   "get-service-provider-config <destination>",
	Short: "Get the default service provider configuration",
	Long: `Retrieve the default service provider configuration in JSON format. 
It will generate a "service_provider_config.json" within the chosen destination path.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)
		if err == nil {
			dest := args[0]
			if !validation.PathExists(dest) {
				// (todo) > use the same error of validator encapsulating this check
				return fmt.Errorf("not a path")
			}
			return nil
		}
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		if config.Values.Debug {
			fmt.Fprintln(os.Stdout, "Generating config ...")
		}
		// (todo) > impl
		fmt.Println("NOT IMPLEMENTED YET.", args)
	},
}
