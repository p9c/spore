package spore

import (
	. "github.com/l0k18/spore/pkg/log"
	"os"
)

type CLI struct {
	*Shell
}

func NewCLI(s *Shell) *CLI {
	return &CLI{Shell: s}
}

func (c *CLI) Run() int {
	Debug("running", os.Args[1:])
	// first we need to check we have go in the shell environment, fetch and unpack it otherwise, write config setting
	// which url and hash of package, write command to set new URL and hash for Go distribution.
	// check if we already have package and the version is current
	// storage for packages - path and then git commit hash
	// prune binary cache. keep whole git repo tied to package, and refer to cache if base is requested again
	// build cache management command - set version to build and use by git commit or tag
	// `go build github.com/org/repo`
	// move binary to binary cache
	// run binary with given parameters
	return 0
}
