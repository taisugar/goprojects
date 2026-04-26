package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use: "remove",
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := strconv.Atoi(args[0])

		initService(formatJSON)
		svc.Remove(i)
		svc.Save()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
