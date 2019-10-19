package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var scanner *bufio.Scanner

	// command line error checking
	switch len(os.Args) {
	case 2: // use input file
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("failed to open %s", os.Args[1])
		}

		scanner = bufio.NewScanner(file)
	case 1: // use standard input
		scanner = bufio.NewScanner(os.Stdin)
	default: // wrong number of arguments
		log.Fatalf("Usage: %s [ filename ]", os.Args[0])
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	lineCount := 0
	for scanner.Scan() {
		lineCount++
		fmt.Println(scanner.Text())
	}
	fmt.Println(lineCount)
}
