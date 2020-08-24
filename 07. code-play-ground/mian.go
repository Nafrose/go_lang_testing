package main

import (
	"fmt"
)

type IWriter interface {
	Write(line string)
}

type Writer struct {
	IsError bool
	IWriter
}

type DefaultLogWriter struct {
	IsError bool
}

func (d DefaultLogWriter) Write(line string) {
	fmt.Println(line)
}

func main() {
	var writer Writer = Writer{IsError: true, IWriter: DefaultLogWriter{}}
	writer.Write("me hello")
	fmt.Println("Hello, playground")
}
