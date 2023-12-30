package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blurhash",
	Short: "Blurhash command line tool",
	Long:  "Blurhash command line tool",
	Example: `
	blurhash encode --filepath ./smile.png --xcomponent 4 --ycomponent 3
	`,
	Version: "v1",
	Run:     root,
}

func root(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func Execute() error {
	return rootCmd.Execute()
}
