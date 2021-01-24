package util

import (
	"crypto/sha256"
	"encoding/hex"
	. "github.com/l0k18/sporeOS/pkg/log"
	"hash"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DownloadFile will download a url to a local file. It's efficient because it
// will write as it downloads and not load the whole file into memory.
func DownloadFile(directory string, url string, hashS string) (writtenFileName string, err error) {
	splitURL := strings.Split(url, "/")
	writtenFileName = filepath.Join(directory, splitURL[len(splitURL)-1])
	// check if the file exists and compute its' checksum - if it has the right checksum no need to download
	if FileExists(writtenFileName) {
		var hasher hash.Hash
		hasher = sha256.New()
		var f *os.File
		if f, err = os.Open(writtenFileName); !Check(err) {
			if _, err = io.Copy(hasher, f); !Check(err) {
				if hex.EncodeToString(hasher.Sum(nil)) == hashS {
					// no need to download it as it is already correct
					Debug("file already downloaded")
					return
				}
			}
		}
	}
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
	buf := make([]byte, 8192)
	// tot := 0
out:
	for {
		var n int
		var rerr error
		if n, rerr = resp.Body.Read(buf); Check(err) {
		}
		// tot += n
		// Debug("read", n, "of", tot, "bytes from", url)
		if _, err = out.Write(buf[:n]); Check(err) {
		}
		// Debug("wrote", n, "bytes of", tot, "to", writtenFileName)
		if rerr == io.EOF {
			break out
		}
	}
	// if _, err = io.Copy(out, resp.Body); Check(err) {
	// }
	return
}
