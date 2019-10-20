package tokens

import (
	"errors"
	"strings"
)

// Tokens contains comment tokens
type Tokens struct {
	Inline     string
	BlockLeft  string
	BlockRight string
}

// Style identifies general groups of comments for languages
type Style int

const (
	CStyle Style = iota // C/C++ style
	Python
	Shell
)

// map of tokens to its lexeme
var tokens = map[Style]Tokens{
	CStyle: Tokens{"//", "/*", "*/"},
	Python: Tokens{"#", "'''", "'''"},
	Shell:  Tokens{"#", "<#", "#>"},
}

// NewTokens - returns comment tokens of language identified using file extension
func NewTokens(f string) (Tokens, error) {
	style, err := GetStyle(f)
	if err != nil {
		return Tokens{}, err
	}

	return tokens[style], nil
}

// GetStyle returns proper commenting style from file ext
func GetStyle(f string) (Style, error) {
	r := strings.Split(f, ".")

	// get extension of file
	ext := r[len(r)-1]

	switch ext {
	case "cc", "h", "c", "C", "c++", "cpp", "java":
		return CStyle, nil
	case "py":
		return Python, nil
	case "sh":
		return Shell, nil
	default:
		return 0, errors.New("invalid file extension")
	}
}
