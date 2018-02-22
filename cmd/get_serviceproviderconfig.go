package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/defaults"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/cobra"
)

func init() {
	scimd.AddCommand(getServiceProviderConfig)
}

const getServiceProviderConfigArgName = "destination"

var getServiceProviderConfig = &cobra.Command{
	Use:   fmt.Sprintf("get-service-provider-config <%s>", getServiceProviderConfigArgName),
	Short: "Get the default service provider configuration",
	Long: `Retrieve the default service provider configuration in JSON format. 
It will generate a "service_provider_config.json" within the chosen destination path.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)
		if err == nil {
			dest := args[0]
			errs := validation.Validator.Var(dest, "pathexists,isdir")
			if errs != nil {
				return fmt.Errorf("%s%s", getServiceProviderConfigArgName, validation.Errors(errs))
			}

			return nil
		}
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		if config.Values.Debug {
			fmt.Fprintln(os.Stdout, "Generating JSON ...")
		}

		bytes, err := json.MarshalIndent(defaults.ServiceProviderConfig, "", "  ")
		check(err)

		dest := filepath.Join(args[0], "service_provider_config.json")
		file, err := os.Create(dest)
		defer file.Close()
		check(err)

		if config.Values.Debug {
			fmt.Fprintf(os.Stdout, "Writing JSON at \"%s\" ...\n", dest)
		}

		_, err = file.Write(bytes)
		check(err)

		if config.Values.Debug {
			fmt.Fprintln(os.Stdout, "Done")
		}
	},
}

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
