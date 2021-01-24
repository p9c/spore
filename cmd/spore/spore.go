package spore

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	
	. "github.com/l0k18/sporeOS/pkg/log"
	"github.com/l0k18/sporeOS/pkg/util"
)

type Shell struct {
	dataDir string
	config  string
}

func New() *Shell {
	dd := util.Dir("spore", false)
	s := &Shell{
		dataDir: dd,
		config:  filepath.Join(dd, "spore.json"),
	}
	util.EnsureDir(s.config)
	var err error
	Debug(runtime.GOOS, runtime.GOARCH)
	var wf string
	if info, ok := goVersions[runtime.GOOS][runtime.GOARCH]; ok {
		Debug("downloading go", info)
		if wf, err = util.DownloadFile(s.dataDir, info.url, info.hash); Check(err) {
		}
		Debug("download completed", wf)
	}
	gopath := filepath.Join(s.dataDir, "go")
	if strings.HasSuffix(wf, ".tar.gz") {
		// unpack the archive if it isn't already
		if !util.FileExists(gopath) {
			Debug("unpacking archive")
			var r *os.File
			if r, err = os.Open(wf); !Check(err) {
				ExtractTarGz(r, s.dataDir)
			}
		}
	} else if strings.HasSuffix(wf, ".zip") {
		if !util.FileExists(gopath) {
			Debug("unpacking archive")
			if _, err = Unzip(wf, s.dataDir); Check(err) {
			}
		}
	}
	return s
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {
	var filenames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()
	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}
		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		_, err = io.Copy(outFile, rc)
		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}
		Debug("unpacked", fpath)
	}
	return filenames, nil
}

func ExtractTarGz(gzipStream io.Reader, prefix string) {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		Fatal("ExtractTarGz: NewReader failed")
	}
	tarReader := tar.NewReader(uncompressedStream)
	var header *tar.Header
out:
	for {
		header, err = tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
			break
		}
		switch header.Typeflag {
		case tar.TypeDir:
			// no need to worry about these, directories are made for files
		case tar.TypeReg:
			var outFile *os.File
			fp := filepath.Join(prefix, header.Name)
			util.EnsureDir(fp)
			outFile, err = os.Create(fp)
			if err != nil {
				Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
				break out
			}
			if _, err = io.Copy(outFile, tarReader); err != nil {
				Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
				break out
			}
			if err = outFile.Close(); Check(err) {
				break out
			}
			Debug("unpacked", fp)
		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %v in %s",
				header.Typeflag,
				header.Name,
			)
		}
		
	}
}

type downloadInfo struct {
	url, hash string
}

var goVersions = map[string]map[string]downloadInfo{
	"darwin": {
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.darwin-amd64.tar.gz",
			"50a64d6a7ef85510321f0cbcd64e7c72f7e82e27c22f0ba475b9b6b6213f136e",
		},
	},
	"linux": {
		"386": downloadInfo{
			"https://golang.org/dl/go1.14.14.linux-386.tar.gz",
			"b08e088ba99134035782c71aeaf139f36d2306eb88eddc22c1278b8b446f157e",
		},
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.linux-amd64.tar.gz",
			"6f1354c9040d65d1622b451f43c324c1e5197aa9242d00c5a117d0e2625f3e0d",
		},
		"arm64": downloadInfo{
			"https://golang.org/dl/go1.14.14.linux-arm64.tar.gz",
			"511d764197121f212d130724afb9c296f0cb4a22424e5ae956a5cc043b0f4a29",
		},
		"armv6l": downloadInfo{
			"https://golang.org/dl/go1.15.7.linux-armv6l.tar.gz",
			"e4d614c23b77a367becaeac3032cf4911793363a33efa299d29440be3d66234b",
		},
	},
	"windows": {
		"386": downloadInfo{
			"https://golang.org/dl/go1.14.14.windows-386.zip",
			"60ebb9f44549f4827bd29bab822ad881cec6d0f83fff49bda7ad20e69b7b4e7b",
		},
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.windows-amd64.zip",
			"88e6be798902d802481b83015e23f6e587cbe0e58766dfa7959d1032865f6bab",
		},
	},
	"freebsd": {
		"386": downloadInfo{
			"https://golang.org/dl/go1.14.14.freebsd-386.tar.gz",
			"7865dffe01499e5e26a40ebc15e068e683e64a2f2edff7440fc9802b02f122bb",
		},
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.freebsd-amd64.tar.gz",
			"a4fab9549523eefe4cdb4d1334144cb51825db2cfe7993497773f5c9349f6647",
		},
	},
}

/*
go1.14.14.darwin-amd64.tar.gz 	Archive 	macOS 	x86-64 	120MB 	50a64d6a7ef85510321f0cbcd64e7c72f7e82e27c22f0ba475b9b6b6213f136e
go1.14.14.linux-386.tar.gz 	Archive 	Linux 	x86 	100MB 	b08e088ba99134035782c71aeaf139f36d2306eb88eddc22c1278b8b446f157e
go1.14.14.linux-amd64.tar.gz 	Archive 	Linux 	x86-64 	118MB 	6f1354c9040d65d1622b451f43c324c1e5197aa9242d00c5a117d0e2625f3e0d
go1.14.14.linux-arm64.tar.gz 	Archive 	Linux 	ARMv8 	97MB 	511d764197121f212d130724afb9c296f0cb4a22424e5ae956a5cc043b0f4a29
go1.14.14.linux-armv6l.tar.gz 	Archive 	Linux 	ARMv6 	97MB 	e4d614c23b77a367becaeac3032cf4911793363a33efa299d29440be3d66234b
go1.14.14.windows-386.zip 	Archive 	Windows 	x86 	113MB 	60ebb9f44549f4827bd29bab822ad881cec6d0f83fff49bda7ad20e69b7b4e7b
go1.14.14.windows-amd64.zip 	Archive 	Windows 	x86-64 	132MB 	88e6be798902d802481b83015e23f6e587cbe0e58766dfa7959d1032865f6bab
go1.14.14.freebsd-386.tar.gz 	Archive 	FreeBSD 	x86 	100MB 	7865dffe01499e5e26a40ebc15e068e683e64a2f2edff7440fc9802b02f122bb
go1.14.14.freebsd-amd64.tar.gz 	Archive 	FreeBSD 	x86-64 	118MB 	a4fab9549523eefe4cdb4d1334144cb51825db2cfe7993497773f5c9349f6647

*/
