package command

import (
	"fmt"

	"github.com/hadihammurabi/git-semver/service"
	"github.com/hadihammurabi/go-ioc/ioc"
	"github.com/spf13/cobra"
)

func get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get current version",
		Args:  cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
			tags := ioc.Get(service.GitService{}).GetValidTags()
			if len(tags) > 0 {
				fmt.Println(tags[len(tags)-1])
			}
		},
	}

	return cmd
}
