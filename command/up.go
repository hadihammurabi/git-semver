package command

import (
	"fmt"

	"github.com/hadihammurabi/git-semver/service"
	"github.com/hadihammurabi/go-ioc/ioc"
	"github.com/spf13/cobra"
)

func up() *cobra.Command {
	var major, minor, patch, tag bool

	cmd := &cobra.Command{
		Use:   "up",
		Short: "move current version to next version",
		Args:  cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
			tags := ioc.Get(service.GitService{}).GetValidTags()
			latest := "0.0.0"
			if len(tags) > 0 {
				latest = tags[len(tags)-1]
			}

			semverUp := service.SemverUpPatch
			if major {
				semverUp = service.SemverUpMajor
			} else if minor {
				semverUp = service.SemverUpMinor
			}

			newVer := ioc.Get(service.SemverService{}).Up(latest, semverUp)
			fmt.Println(newVer)

			if tag {
				ioc.Get(service.GitService{}).CreateTag(newVer)
			}
		},
	}

	cmd.Flags().BoolVarP(&major, "major", "M", false, "view next major version")
	cmd.Flags().BoolVarP(&minor, "minor", "m", false, "view next minor version")
	cmd.Flags().BoolVarP(&patch, "patch", "p", false, "view next patch version")
	cmd.Flags().BoolVarP(&tag, "tag", "t", false, "update version and create new tag")

	return cmd
}
