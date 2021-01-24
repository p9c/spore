package main

var Commands = map[string][]string{
	"install": {
		`go install -v %ldflags .`,
	},
	"run": {
		`go install -v %ldflags .`,
		`reset`,
		`sporeOS`,
	},
	"hello": {
		`go install -v %ldflags .`,
		`reset`,
		`sporeOS github.com/l0k18/sporeOS/cmd/hello test`,
	},
	"windows": {
		`go build -v -ldflags="-H windowsgui" %ldflags"`,
	},
	"update": {
		"go install -v %ldflags ./stroy/.",
	},
}
