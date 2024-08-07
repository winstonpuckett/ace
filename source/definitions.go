package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Definition interface {
	Run()
	String() string
}

type String struct {
	value string
}

func (s String) Run() {
	myStack.Push(s)
}

func (s String) String() string {
	return s.value
}

type Script struct {
	Parts []string
}

func (s Script) Run() {
	stackValues := make([]string, 0, myStack.Len())
	for _, part := range s.Parts {
		if len(part) < 1 {
			continue
		} else if part[0] == '+' {
			stackNumber, err := strconv.Atoi(part[1:])
			if err != nil {
				panic(err)
			}

			for stackNumber > len(stackValues) {
				stackValues = append(stackValues, myStack.Pop().(Definition).String())
			}

			stringBuilder.WriteString(stackValues[stackNumber-1])
		} else {
			stringBuilder.WriteString(part)
		}
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", stringBuilder.String())
	} else {
		cmd = exec.Command("sh", "-c", stringBuilder.String())
	}

	stringBuilder.Reset()

	output, err := cmd.Output()

	if len(output) > 0 {
		s := string(output)
		myStack.Push(String{value: s})
		os.Stdout.WriteString(s)
	}

	if err != nil && len(err.Error()) > 0 {
		s := err.Error()
		myStack.Push(String{value: s})
		os.Stderr.WriteString(s)
	}
}

func (s Script) String() string {
	var parts []string
	for idx, el := range s.Parts {
		if len(el) < 1 {
			continue
		}
		parts = append(parts, string("\n\tsub ")+fmt.Sprint(idx)+el)
	}

	return fmt.Sprintf("Script: %s", strings.Join(parts, ", "))
}

type Word struct {
	Key         string
	Definitions []Definition
}

func (w Word) Run() {
	definitions[w.Key] = w.Definitions
}

func (w Word) String() string {
	var definitions []string
	for idx, el := range w.Definitions {
		definitions = append(definitions, string("\n\tsub")+fmt.Sprint(idx)+el.String())
	}

	return fmt.Sprintf("Word: %s with definition, %s", string(w.Key), strings.Join(definitions, ", "))
}

type EnvironmentVariable struct {
	Name string
}

func (e EnvironmentVariable) Run() {
	variable := os.Getenv(e.Name)
	myStack.Push(variable)
}

func (e EnvironmentVariable) String() string {
	return fmt.Sprintf("EnvironmentVariable: %s", string(e.Name))
}

type Reference struct {
	Name string
}

func (r Reference) Run() {
	target, ok := definitions[r.Name]

	if !ok {
		panic("could not find key: " + r.Name)
	}

	for _, el := range target {
		el.Run()
	}
}

func (r Reference) String() string {
	return fmt.Sprintf("Reference: %s", string(r.Name))
}
