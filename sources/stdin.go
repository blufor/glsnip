package sources

import (
	"io"
	"os"
)

type Pipe struct {
	d *os.File
}

func NewPipe() (f *Pipe) {
	f = new(Pipe)
	f.d = os.Stdin
	return
}

func (f *Pipe) Content() (c string, err error) {
	b, err := io.ReadAll(f.d)
	if err != nil {
		return
	}
	c = string(b)
	return
}

func (f *Pipe) Path() string {
	return "STDIN"
}

func (f *Pipe) Size() (s int64) {
	fs, err := f.d.Stat()
	if err != nil {
		return
	}
	s = fs.Size()
	return
}

func (f *Pipe) Close() (err error) {
	return nil
}
