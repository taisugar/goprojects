package adapter_cli

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func (c *CLI) newListCmd() *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			tasks := c.svc.List()

			table := table.New(os.Stdout)
			table.SetRowLines(false)
			table.SetHeaders("#", "Title", "Description", "Created At", "Completed", "Completed At")

			for i, task := range tasks {
				completed := "❌"
				completedAt := ""
				if task.IsCompleted {
					completed = "✅"
				}
				if task.CompletedAt != nil {
					completedAt = timediff.TimeDiff(*task.CompletedAt)
				}
				table.AddRow(strconv.Itoa(i), task.Title, task.Description, task.CreatedAt.Format(time.RFC1123), completed, completedAt)
			}
			table.Render()
		},
	}
}
