package spore

import (
	"github.com/l0k18/spore/pkg/util"
)

type Shell struct {
	dataDir string
}

func New() *Shell {
	s := &Shell{
		dataDir: util.Dir("spore", false),
	}
	return s
}
