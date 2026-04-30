package adapter_cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *CLI) newCompleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "complete [index]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			err = c.svc.Complete(i)
			if err != nil {
				return err
			}

			fmt.Println("Task marked as completed")
			return nil
		},
	}
}
