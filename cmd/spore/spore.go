package spore

import "github.com/l0k18/OSaaS/pkg/appdata"

type Shell struct {
	dataDir string
}

func New() *Shell {
	s := &Shell{
		dataDir: appdata.Dir("spore", false),
	}
	return s
}