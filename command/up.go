package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func up() *cobra.Command {
	var major, minor, patch, tag bool

	cmd := &cobra.Command{
		Use:   "up",
		Short: "move current version to next version",
		Args:  cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(major, minor, patch, tag)
		},
	}

	cmd.Flags().BoolVarP(&major, "major", "M", false, "view next major version")
	cmd.Flags().BoolVarP(&minor, "minor", "m", false, "view next minor version")
	cmd.Flags().BoolVarP(&patch, "patch", "p", false, "view next patch version")
	cmd.Flags().BoolVarP(&tag, "tag", "t", false, "update version and create new tag")

	return cmd
}
