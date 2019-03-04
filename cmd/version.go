package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is go-hello-world's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-hello-world v1.0")
	},
}
