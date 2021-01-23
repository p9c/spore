package main

import (
	. "github.com/l0k18/spore/pkg/log"
	"os"
)

func main() {
	Debug("hello", os.Args[1:])
}
