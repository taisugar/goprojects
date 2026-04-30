package adapter_cli

import (
	"task-manager/internal/app"

	"github.com/spf13/cobra"
)

type CLI struct {
	svc *app.TaskService
}

type format string

const (
	formatCSV  format = "csv"
	formatJSON format = "json"
)

func RegisterCLI(svc *app.TaskService) *cobra.Command {
	c := &CLI{svc: svc}

	rootCmd := &cobra.Command{
		Use: "todo",
	}

	rootCmd.AddCommand(
		c.newAddCmd(),
		c.newListCmd(),
		c.newEditCmd(),
		c.newCompleteCmd(),
		c.newRemoveCmd(),
	)

	return rootCmd
}
