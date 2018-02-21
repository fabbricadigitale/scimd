package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

func init() {
	scimd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SCIMD",
	Long:  `Shows the version number of SCIMD.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "Version %s.\n", version)
	},
}
