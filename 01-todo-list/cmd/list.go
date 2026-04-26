package cmd

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		initService(formatJSON)

		table := table.New(os.Stdout)
		table.SetRowLines(false)
		table.SetHeaders("#", "Title", "Description", "Created At", "Completed", "Completed At")

		for i, task := range svc.List() {
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

func init() {
	rootCmd.AddCommand(listCmd)
}
