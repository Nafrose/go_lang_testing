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

}

func (d DefaultLogWriter) Write(line string) {
	fmt.Println(line)
}

func CreateDefaultErrorLogWriter() IWriter {
	write := func (line string){fmt.Println(line)}
	return Writer{IsError: true, IWriter: func Write(line string){fmt.Println(line)}}
}

func main() {
	var writer Writer = CreateDefaultErrorLogWriter()
	writer.Write("me hello")
	fmt.Println("Hello, playground")
}
