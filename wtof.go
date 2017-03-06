package wtof

import (
	"bufio"
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
		// b := make([]byte, bufsize)
		reader := bufio.NewReader(w)
		writer := bufio.NewWriter(iow)
		for {
			b, err := reader.ReadByte()
			// i, err := w.Read(b)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			writer.WriteByte(b)
			// iow.Write(b[:i])

		}
	}()
	return t
}
func (t *Wtof) Close() {
	t.File.Close()
	<-t.done
}
