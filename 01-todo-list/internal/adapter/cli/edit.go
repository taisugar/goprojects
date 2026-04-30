package adapter_cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *CLI) newEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "edit [index]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			title, _ := cmd.Flags().GetString("title")
			desc, _ := cmd.Flags().GetString("desc")

			err = c.svc.Edit(i, title, desc)
			if err != nil {
				return err
			}

			fmt.Println("Task edited")
			return nil
		},
	}

	cmd.Flags().String("title", "", "title")
	cmd.Flags().String("desc", "", "description")

	return cmd
}
