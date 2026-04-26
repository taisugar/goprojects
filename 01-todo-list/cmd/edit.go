package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use: "edit",
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := strconv.Atoi(args[0])
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")

		initService(formatJSON)
		svc.Edit(i, title, desc)
		svc.Save()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().String("title", "", "title")
	editCmd.Flags().String("desc", "", "desc")
}
