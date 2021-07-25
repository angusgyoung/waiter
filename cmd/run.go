package cmd

import (
	"github.com/angusgyoung/waiter/internal"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskName = args[0]
		internal.RunTask(taskName)
	},
}

func init() {
	taskCmd.AddCommand(runCmd)
}
