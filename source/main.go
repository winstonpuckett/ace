package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

var myStack = stack.New()
var definitions = make(map[string][]Definition, 0)

var stringBuilder strings.Builder
var scanner Scanner

func main() {
	if len(os.Args) > 1 {
		Open(os.Args[1])
	}

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		scanner = &StringScanner{
			source:   []byte(in.Text()),
			position: 0,
		}
		ParseAndExecute()
	}

	if in.Err() != nil {
		panic(in.Err().Error())
	}
}
