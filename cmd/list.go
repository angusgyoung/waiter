package cmd

import (
	"fmt"

	"github.com/angusgyoung/waiter/internal"
	"github.com/angusgyoung/waiter/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		var tasks []task.Task = internal.Conf.Tasks

		for _, task := range tasks {
			fmt.Printf("Name %v, command %v", task.Name, task.Command)
		}
	},
}

func init() {
	taskCmd.AddCommand(listCmd)
}
