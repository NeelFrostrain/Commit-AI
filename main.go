package main

import (
	"github.com/NeelFrostrain/Commit-Ai/cmd"
)

var (
	// Version is set via ldflags during build
	Version   = "dev"
	BuildDate = "unknown"
	GitCommit = "unknown"
)

func main() {
	cmd.SetVersion(Version, BuildDate, GitCommit)
	cmd.Execute()
}
