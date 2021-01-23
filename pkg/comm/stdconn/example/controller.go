package main

import (
	"github.com/l0k18/OSaaS/pkg/comm/stdconn/example/hello/hello"
	"github.com/l0k18/OSaaS/pkg/comm/stdconn/worker"
	log "github.com/l0k18/OSaaS/pkg/util/logi"
	qu "github.com/l0k18/OSaaS/pkg/util/quit"
)

func main() {
	log.L.SetLevel("trace", true, "pod")
	Info("starting up example controller")
	cmd, _ := worker.Spawn(qu.T(), "go", "run", "hello/worker.go")
	client := hello.NewClient(cmd.StdConn)
	Info("calling Hello.Say with 'worker'")
	Info("reply:", client.Say("worker"))
	Info("calling Hello.Bye")
	Info("reply:", client.Bye())
	if err := cmd.Kill(); Check(err) {
	}
}
