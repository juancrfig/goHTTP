package utils

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

func PrintLinesFromFile(r io.Reader, w io.Writer) error {

	if w == nil {
		w = os.Stdout
	}

	var readingBuf []byte = make([]byte, 8)  
	var leftover *bytes.Buffer = &bytes.Buffer{}

	for {
		n, err := r.Read(readingBuf)

		if err != io.EOF && err != nil {
			return err
		}

		if n < 1 {
			fmt.Println("")
			return nil
		}

		data := readingBuf[:n]

		//fmt.Printf("Leftover before if: %s", leftover.String())
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			leftover.WriteString(string(data[:i]))
			fmt.Fprintf(w, "read: %s\n", leftover.String())

			leftover.Reset()
			leftover.WriteString(string(data[i+1:]))
			continue
		}
		leftover.WriteString(string(data))
	}
}
