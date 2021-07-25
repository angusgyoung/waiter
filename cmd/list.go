package cmd

import (
	"os"

	"github.com/angusgyoung/waiter/internal"
	"github.com/angusgyoung/waiter/internal/task"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",

	Run: func(cmd *cobra.Command, args []string) {
		internal.Log.Debug("Listing tasks")

		var tasks []task.Task = internal.Conf.Tasks

		if len(tasks) > 0 {
			t := table.NewWriter()
			t.SetStyle(table.StyleLight)
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Name", "Command"})

			for i, task := range tasks {
				t.AppendRow([]interface{}{i, task.Name, task.Command})
			}

			t.Render()
		}

	},
}

func init() {
	taskCmd.AddCommand(listCmd)
}
