package adapter_cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *CLI) newRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "remove [index]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			err = c.svc.Remove(i)
			if err != nil {
				return err
			}

			fmt.Println("Task removed")
			return nil
		},
	}
}
