package util

import (
	. "github.com/l0k18/spore/pkg/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(directory string, url string) (writtenFileName string, err error) {
	splitURL := strings.Split(url, "/")
	writtenFileName = filepath.Join(directory, splitURL[len(splitURL)-1])
	// Get the data
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return
	}
	defer func() {
		if err = resp.Body.Close(); Check(err) {
		}
	}()
	var out *os.File
	if out, err = os.Create(writtenFileName); Check(err) {
		return
	}
	// Create the file
	defer func() {
		if err = out.Close(); Check(err) {
		}
	}()
	// Write the body to file
	if _, err = io.Copy(out, resp.Body); Check(err) {
	}
	return
}
