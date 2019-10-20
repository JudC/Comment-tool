package scanner

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/JudC/Comment-tool/pkg/tokens"
)

// CommentScanner - contains state of scanning
type CommentScanner struct {
	tokens      tokens.Tokens // comment tokens for language of file
	fileName    string        // name of file opened
	isMulti     bool          // whether current line scanned is part of block comment
	SingleCount int           // # of single line comments
	MultiCount  int           // # of multi-line comments
	BlockCount  int           // # of blocks
	TodoCount   int           // # of todos
	TotalCount  int           // # of comments in total
}

// NewCommentScanner - create a comment scanner for a particular file
func NewCommentScanner(f string) CommentScanner {
	tokens, err := tokens.NewTokens(f)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	return CommentScanner{tokens: tokens, fileName: f}
}

func (s *CommentScanner) newScanner() *bufio.Scanner {
	f, err := os.Open(s.fileName)
	if err != nil {
		log.Fatalf("failed to open %s", s.fileName)
	}

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner
}

// GetLineCount returns the total lines. Includes white space.
func (s *CommentScanner) GetLineCount() int {
	scanner := s.newScanner()

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

// GetCommentCount determines comment counts
func (s *CommentScanner) GetCommentCount() {
	scanner := s.newScanner()

	for scanner.Scan() {
		line := scanner.Text()

		single := 0
		multi := 0
		todo := 0

		if !s.isMulti {
			single, multi, todo = s.getCommentCountFromLine(line)

			if multi > 0 {
				s.BlockCount++
			}

			if single > 0 || multi > 0 {
				s.TotalCount++
			}
		} else {
			multiEndIndex := strings.Index(line, s.tokens.BlockRight)

			if multiEndIndex != -1 {
				//get number of TODOs before end of block
				todo = s.getTODOCountFromLine(line[:multiEndIndex])
				s.TodoCount += todo

				s.isMulti = false
				if multiEndIndex < len(line)-1 {
					single, multi, todo = s.getCommentCountFromLine(line[multiEndIndex+1:])

				}
			}
			s.MultiCount++

		}

		s.SingleCount += single
		s.MultiCount += multi
		s.TodoCount += todo
	}
}

// returns single comments count, multi
func (s *CommentScanner) getCommentCountFromLine(line string) (int, int, int) {
	singleIndex := strings.Index(line, s.tokens.Inline)
	multiStartIndex := strings.Index(line, s.tokens.BlockLeft)

	// single line comment
	if singleIndex != -1 && (singleIndex < multiStartIndex || multiStartIndex == -1) {
		todo := s.getTODOCountFromLine(line[singleIndex:])
		return 1, 0, todo
	} else if multiStartIndex != -1 { // multiline comment
		s.isMulti = true
		todo := s.getTODOCountFromLine(line[multiStartIndex:])
		return 0, 1, todo
	}

	return 0, 0, 0
}

// finds the number of TODOs withina commented line
// TODO in this context is defined to be any occurence of "TODO" that
// is not preceded by an alpha-numeric character
func (s *CommentScanner) getTODOCountFromLine(line string) int {
	count := 0
	remainingLine := line
	for {
		i := strings.Index(remainingLine, "TODO")
		if i == -1 {
			break
		}

		if i == 0 {
			count++
		} else {
			c, _ := utf8.DecodeRuneInString(string(remainingLine[i-1]))

			if unicode.IsSpace(c) || (!unicode.IsDigit(c) && !unicode.IsLetter(c)) {
				count++
			}
		}
		remainingLine = remainingLine[i+len("TODO"):]
	}

	return count
}
