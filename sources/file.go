package sources

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type File struct {
	d *os.File
	s fs.FileInfo
	p string
}

func NewFile(path string) (f *File, err error) {
	f = new(File)

	f.p = filepath.Base(path)

	f.d, err = os.Open(path)
	if err != nil {
		return
	}

	f.s, err = f.d.Stat()
	return
}

func (f *File) Content() (c string, err error) {
	b, err := io.ReadAll(f.d)
	if err != nil {
		return
	}
	c = string(b)
	return
}

func (f *File) Path() string {
	return f.p
}

func (f *File) Size() int64 {
	return f.s.Size()
}

func (f *File) Close() (err error) {
	return f.d.Close()
}
