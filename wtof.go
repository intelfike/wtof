package wtof

import (
	"fmt"
	"io"
	"os"
)

type Wtof struct {
	done chan interface{}
	File *os.File
}

func New(iow io.Writer, bufsize int) *Wtof {
	t := new(Wtof)
	t.done = make(chan interface{})
	var w *os.File
	w, t.File, _ = os.Pipe()
	go func() {
		defer func() {
			w.Close()
			t.done <- nil
		}()
		b := make([]byte, bufsize)
		for {
			i, err := w.Read(b)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			iow.Write(b[:i])
		}
	}()
	return t
}
func (t *Wtof) Close() {
	t.File.Close()
	<-t.done
}
