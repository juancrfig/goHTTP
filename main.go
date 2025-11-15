package main

import (
	"log"
	"os"
	
	"github.com/juancrfig/goHTTP/utils"
)


func main() {

	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error while opening the file: %v\n", err)
	}
	defer file.Close()

	err = utils.PrintLinesFromFile(file, nil)

	if err != nil {
		log.Fatalf("Error while getting lines string: %v\n", err)
	}
}
