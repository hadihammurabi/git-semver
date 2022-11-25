package main

import (
	"os"

	"github.com/hadihammurabi/git-semver/command"
	"github.com/hadihammurabi/git-semver/service"
)

func init() {
	service.Build()
}

func main() {
	cmd := command.Build()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
