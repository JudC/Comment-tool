package main

import (
	"fmt"
	"log"
	"os"

	sc "github.com/JudC/Comment-tool/pkg/scanner"
)

func main() {
	var fileName string
	// command line error checking
	switch len(os.Args) {
	case 2: // use input file
		fileName = os.Args[1]
	default: // wrong number of arguments
		log.Fatalf("Usage: %s [ filename ]", os.Args[0])
	}

	// get instance of file scanner
	cs := sc.NewCommentScanner(fileName)

	// print total number of lines
	fmt.Printf("Total # of lines: %v\n", cs.GetLineCount())

	// parse comments
	cs.GetCommentCount()

	// print results of counting single and multi-line comments
	fmt.Printf("Total # of comment lines: %v\nTotal # of single line comments: %v\n"+
		"Total # of comment lines within block comments: %v\nTotal # of block line comments: %v\n"+
		"Total # of TODO's: %v", cs.TotalCount, cs.SingleCount, cs.MultiCount, cs.BlockCount, cs.TodoCount)
}
