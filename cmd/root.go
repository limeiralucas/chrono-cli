package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Short: "Time converter",
}

func Execute() error {
	rootCmd.AddCommand(togglCmd)

	return rootCmd.Execute()
}
