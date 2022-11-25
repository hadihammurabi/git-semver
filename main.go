package main

import (
	"os"

	"github.com/hadihammurabi/git-semver/command"
)

func main() {
	cmd := command.Build()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
