package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use: "complete",
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := strconv.Atoi(args[0])

		initService(formatJSON)
		svc.Complete(i)
		svc.Save()
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
