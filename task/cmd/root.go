package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
	},
}
