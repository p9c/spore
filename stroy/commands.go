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
	"hello": {
		`go install -v %ldflags .`,
		`reset`,
		`spore github.com/l0k18/spore/cmd/hello test`,
	},
	"windows": {
		`go build -v -ldflags="-H windowsgui" %ldflags"`,
	},
	"update": {
		"go install -v %ldflags ./stroy/.",
	},
}
