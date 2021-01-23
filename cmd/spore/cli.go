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
	// start by fetching the package directory
	
	return 0
}
