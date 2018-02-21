package cmd

import (
	"fmt"
	"os"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/cobra"
)

func init() {
	scimd.AddCommand(getConfig)
}

var getConfig = &cobra.Command{
	Use:   "get-config <destination>",
	Short: "Get the default configuration",
	Long: `Retrieve the default configurations.
It will generate the JSON files representing the default schemas and resource types, within the chosen destination path.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)
		if err == nil {
			dest := args[0]
			errs := validation.Validator.Var(dest, "pathexists,isdir")
			if errs != nil {
				return fmt.Errorf("%s%s", arg, validation.Errors(errs))
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
