package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error while opening the file: %v\n", err)
	}

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(ch)

		leftover := &bytes.Buffer{}
		for {
			buf := make([]byte, 8)
			n, err := f.Read(buf)

			buf = buf[:n]
			if i := bytes.IndexByte(buf, '\n'); i != -1 {
				leftover.WriteString(string(buf[:i]))
				ch <- leftover.String()

				leftover.Reset()

				leftover.WriteString(string(buf[i+1:]))
				continue

			}
			//fmt.Printf("Writing '%s' to leftover: %s\n", string(buf), leftover.String())
			leftover.WriteString(string(buf))

			if err != nil {
				break
			}
		}

	}()

	return ch
}
