package main

import (
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

var myStack = stack.New()
var definitions = make(map[string][]Definition, 0)

var stringBuilder strings.Builder
var scanner Scanner

func main() {
	Open(os.Args[1])
}
