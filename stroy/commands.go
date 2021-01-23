package main

var Commands = map[string][]string{
	"install": {
		`go install -v %ldflags .`,
	},
	"run": {
		`go install -v %ldflags .`,
		`reset`,
		`spore`,
	},
	"windows": {
		`go build -v -ldflags="-H windowsgui" %ldflags"`,
	},
	"update": {
		"go install -v %ldflags ./stroy/.",
	},
}
