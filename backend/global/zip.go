package global

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// MakeZipfile zip파일로 압축
// filename: 압축할 파일명
// name: 압축할 파일의 이름 목록(한글 유효)
// files: 압축할 파일의 전체경로 목록
func MakeZipfile(zfile string, fname []string, files []string) {

	fmt.Printf("\n\nMakeZipfile:\n  %+v\n  %+v\n  %+v\n", zfile, fname, files)

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	f1 := filepath.Join("webdata", zfile)
	fmt.Printf("MakeZipfile:zipfile  %+v\n", f1)

	buf, e := os.OpenFile(f1, flags, 0644)
	if e != nil {
		log.Fatalf("Failed to open zip for writing: %s", e)
	}

	defer buf.Close()

	zif := zip.NewWriter(buf)
	defer zif.Close()

	for i, f := range files {
		h := &zip.FileHeader{Name: fname[i], Method: zip.Deflate, Flags: 0x800}
		zf, e := zif.CreateHeader(h)
		if e != nil {
			log.Fatal(e)
			continue
		}

		//f2 := filepath.Join("webdata", f)
		f2 := f
		fs, _ := os.Open(f2)
		defer fs.Close()

		io.Copy(zf, fs)
		zif.Flush()
	}

	zif.Close()
}
