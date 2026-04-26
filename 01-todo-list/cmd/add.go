package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")

		initService(formatJSON)
		svc.Add(title, desc)
		svc.Save()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("title", "", "title")
	addCmd.Flags().String("desc", "", "desc")
}
