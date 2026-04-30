package adapter_cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *CLI) newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "add [title]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			title := args[0]
			if title == "" {
				title, _ = cmd.Flags().GetString("title")
			}

			desc, _ := cmd.Flags().GetString("desc")

			c.svc.Add(title, desc)

			fmt.Println("Task added successfully")
			return nil
		},
	}

	cmd.Flags().String("title", "", "title")
	cmd.Flags().String("desc", "", "description")

	return cmd
}
