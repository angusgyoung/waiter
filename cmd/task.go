package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
	Long:  `View, create and run tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
}
