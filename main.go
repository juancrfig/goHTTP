package main

import (
	"log"
	"os"
	"io"
	"fmt"
)


func main() {
	// Basic debugging config 
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error while opening the file: %v", err)
	}
	defer file.Close()

	var buffer = make([]byte, 8)
	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if n < 8 {
			fmt.Printf("read: %s\n", string(buffer[:n]))
		} else {
			fmt.Printf("read: %s\n", string(buffer))
		}
	}
}
