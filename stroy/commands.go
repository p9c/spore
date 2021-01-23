package main

var Commands = map[string][]string{
	"build": {
		"go build -v",
	},
	"windows": {
		`go build -v -ldflags="-H windowsgui \"%ldflags"\"`,
	},
	"": {
		``,
	},
	"stroy": {
		"go install -v %ldflags ./stroy/.",
	},
}
