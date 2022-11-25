package command

import (
	"github.com/spf13/cobra"
)

func get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get current version",
		Args:  cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
