package main

import (
	"log"
	"os"
	
	"github.com/juancrfig/goHTTP/utils"
)


func main() {
	// Basic debugging config 
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error while opening the file: %v", err)
	}
	defer file.Close()

	lines, err := utils.GetLinesFromFile(file)

	if err != nil {
		log.Fatal("Error while getting lines string")
	}

	log.Printf("slice lines: %v", lines)
	
}
