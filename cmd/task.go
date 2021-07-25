package cmd

import (
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
	Long:  `View, create and run tasks`,
}

func init() {
	rootCmd.AddCommand(taskCmd)
}
