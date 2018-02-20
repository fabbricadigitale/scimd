package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func init() {
	scimdCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SCIMD",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version %s.\n", version)
	},
}
