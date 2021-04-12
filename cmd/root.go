package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Long: "dod is a docker container & image deleter",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}