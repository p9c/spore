package main

import (
	"fmt"
	"github.com/l0k18/spore/cmd/spore"
	"github.com/l0k18/spore/pkg/util/logi"
	"github.com/l0k18/spore/version"
	"os"
)

func main() {
	version.URL = URL
	version.GitRef = GitRef
	version.GitCommit = GitCommit
	version.BuildTime = BuildTime
	version.Tag = Tag
	version.Get = GetVersion
	logi.L.SetLevel("debug", false, "spore")
	Debug(version.Get())
	s := spore.New()
	result := s.Main()
	if result != 0 {
		os.Exit(result)
	}
}

var (
	URL       string
	GitRef    string
	GitCommit string
	BuildTime string
	Tag       string
)

func GetVersion() string {
	return fmt.Sprintf(
		`Application Information
	repo: %s
	branch: %s
	commit: %s
	built: %s
	tag: %s
`,
		URL, GitRef, GitCommit, BuildTime, Tag,
	)
}
