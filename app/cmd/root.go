package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker-cli",
		Short: "task tracker",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Welcome to task tracker cli!")
		},
	}

	cmd.AddCommand(newUpdateCommand())
	cmd.AddCommand(newDeleteCommand())
	cmd.AddCommand(newAddCommand())
	cmd.AddCommand(newMarkInProgressCommand())
	cmd.AddCommand(newMarkDoneCommand())
	cmd.AddCommand(newListCommand())

	return cmd
}
