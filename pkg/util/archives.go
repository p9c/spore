package util

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	. "github.com/l0k18/sporeOS/pkg/log"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
			EnsureDir(fp)
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
