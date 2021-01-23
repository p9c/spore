package main

var Commands = map[string][]string{
	"install": {
		`go install -v %ldflags .`,
	},
	"windows": {
		`go build -v -ldflags="-H windowsgui" %ldflags"`,
	},
	"stroy": {
		"go install -v %ldflags ./stroy/.",
	},
}
