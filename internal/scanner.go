package internal

type commentScanner struct {
	fileName            string // name of file
	line                int    // total number of lines of file
	totalComments       int    // total number of comments
	totalSingleComments int    // total number of single line comments
	totalMultiComments  int    // total number of multi line comments
}
