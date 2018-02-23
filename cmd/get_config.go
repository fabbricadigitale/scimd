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
	scimd.AddCommand(getConfig)
}

const getConfigArgName = "destination"

var getConfig = &cobra.Command{
	Use:   fmt.Sprintf("get-config <%s>", getConfigArgName),
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
				return fmt.Errorf("%s%s", getConfigArgName, validation.Errors(errs))
			}

			return nil
		}
		return err
	},
	// Since we do not use nor care configuration values here
	// We do not check and print configuration validation errors at pre-run time
	Run: func(cmd *cobra.Command, args []string) {
		if config.Values.Debug {
			fmt.Fprintln(os.Stdout, "Generating config ...")
		}

		destin := args[0]
		pathSc := filepath.Join(destin, "schemas")
		pathRt := filepath.Join(destin, "resources")

		var e error
		e = os.MkdirAll(pathSc, os.ModePerm)
		check(e)

		groupSc, e := json.MarshalIndent(defaults.GroupSchema, "", "  ")
		check(e)

		userSc, e := json.MarshalIndent(defaults.UserSchema, "", "  ")
		check(e)

		e = os.MkdirAll(pathRt, os.ModePerm)
		check(e)

		groupRt, e := json.MarshalIndent(defaults.GroupResourceType, "", "  ")
		check(e)

		userRt, e := json.MarshalIndent(defaults.UserResourceType, "", "  ")
		check(e)

		fileGroupSc, e := os.Create(filepath.Join(pathSc, "group.json"))
		defer fileGroupSc.Close()
		check(e)

		fileUserSc, e := os.Create(filepath.Join(pathSc, "user.json"))
		defer fileUserSc.Close()
		check(e)

		fileGroupRt, e := os.Create(filepath.Join(pathRt, "group.json"))
		defer fileGroupRt.Close()
		check(e)

		fileUserRt, e := os.Create(filepath.Join(pathRt, "user.json"))
		defer fileUserRt.Close()
		check(e)

		if config.Values.Debug {
			fmt.Fprintf(os.Stdout, "Writing JSON files at \"%s\" ...\n", destin)
		}

		_, e = fileGroupSc.Write(groupSc)
		check(e)

		_, e = fileUserSc.Write(userSc)
		check(e)

		_, e = fileGroupRt.Write(groupRt)
		check(e)

		_, e = fileUserRt.Write(userRt)
		check(e)

		if config.Values.Debug {
			fmt.Fprintln(os.Stdout, "Done.")
		}
	},
	DisableAutoGenTag: true,
}

// (fixme) > default variables contains "meta" and "schema" fields, should not ..
