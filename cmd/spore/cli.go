package spore

type CLI struct {
	*Shell
}

func NewCLI(s *Shell) *CLI {
	return &CLI{Shell: s}
}

func (c *CLI) Run() int {
	return 0
}
