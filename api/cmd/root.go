package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "domane",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(StartCmd())
}
