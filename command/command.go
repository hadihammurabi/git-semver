package command

import (
	"github.com/spf13/cobra"
)

func Build() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "git-semver",
		Short:   "Semantic version utilities with Git integration",
		Example: "git-semver up major",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				cmd.Help()
			}
		},
	}

	cmd.AddCommand(get())
	cmd.AddCommand(up())

	cmd.DisableAutoGenTag = true
	cmd.DisableSuggestions = true

	return cmd
}
