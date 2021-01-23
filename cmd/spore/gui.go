package spore

type GUI struct {
	*Shell
}

func NewGUI(s *Shell) *GUI {
	return &GUI{Shell: s}
}

func (c *GUI) Run() int{
	return 0
}