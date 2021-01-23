package main

import (
	"github.com/l0k18/OSaaS/cmd/spore"
	"os"
)

func main() {
	s := spore.New()
	result := s.Main()
	if result != 0 {
		os.Exit(result)
	}
}
