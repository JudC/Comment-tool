# Comment parser

Determines all comments within a file and prints output in the following format: 

Total # of lines: 61

Total # of comment lines: 19

Total # of single line comments: 9

Total # of comment lines within block comments: 10

Total # of block line comments: 3

Total # of TODO’s: 3

## Environment
Currently uses `go1.12.7` 

## Building

1. Run `go get https://github.com/JudC/Comment-tool`
2. Run `go install github.com/JudC/Comment-tool/...`

## Usage

Use the binary built in the previous step (should be in go/bin)
`command-tool [path-to-file]`

Currently supports python, bash and C-style comments
