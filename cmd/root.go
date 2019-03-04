package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hellocli",
	Short: "go-hello-world cobra test cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("abc")
	},
	Args: cobra.NoArgs,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
