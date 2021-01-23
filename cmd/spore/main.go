package spore

import (
	"os"
)

func (s *Shell) Main() int {
	if len(os.Args) > 1 {
		Debug("running commandline args", os.Args[1:])
		c := NewCLI(s)
		return c.Run()
	} else {
		Debug("launching gui shell")
		g := NewGUI(s)
		return g.Run()
	}
}
