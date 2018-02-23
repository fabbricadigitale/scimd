package cmd

import (
	"fmt"
	"os"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
	branch  string
	summary string
	date    string
)

func init() {
	scimd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of scimd",
	Long: `This command shows the version number of scimd.
	
It also can print detailed information about the build process of the current scimd.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "Version %s.\n", version)

		if config.Values.Debug {
			fmt.Println()
			fmt.Fprintf(os.Stdout, "Commit\t\t%s.\n", commit)
			fmt.Fprintf(os.Stdout, "Branch\t\t%s.\n", branch)
			fmt.Fprintf(os.Stdout, "Summary\t\t%s.\n", summary)
			fmt.Fprintf(os.Stdout, "Build date\t%s.\n", date)
		}
	},
	DisableAutoGenTag: true,
}
