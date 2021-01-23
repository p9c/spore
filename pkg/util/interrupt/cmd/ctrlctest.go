package main

import (
	"fmt"
	
	"github.com/l0k18/OSaaS/pkg/util/interrupt"
)

func main() {
	interrupt.AddHandler(func() {
		fmt.Println("IT'S THE END OF THE WORLD!")
	})
	<-interrupt.HandlersDone
}
