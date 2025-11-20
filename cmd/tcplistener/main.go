package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}

		linesChan := getLinesChannel(conn)

		for line := range linesChan {
			fmt.Println(line)
		}
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
			leftover.WriteString(string(buf))

			if err != nil {
				break
			}
		}

	}()

	return ch
}
