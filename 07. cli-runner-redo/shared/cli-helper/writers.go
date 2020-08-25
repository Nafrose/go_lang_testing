package cliHelper

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

type IWriter interface {
	Write(line string)
}

type Writer struct {
	IWriter
}

type DefaultLogWriter struct{}

func (d DefaultLogWriter) Write(line string) {
	log.Println(line)
}

type DefaultErrorLogWriter struct{}

func (d DefaultErrorLogWriter) Write(line string) {
	log.Fatalln("Error: %s", line)
}

type WritersCollection struct {
	Writers      []Writer
	ErrorWriters []Writer
}

type WritingConfiguration struct {
	IsJsonLog bool
	// if enabled then write to fmt.Fprintf
	IsWriteToHttp bool
	// https://github.com/uber-go/zap
	IsLogUsingZap        bool
	IsWriteToFile        bool
	IsWriteToDirectory   bool
	WriteToFileDirectory string
	WriteToFilePath      string
}

var zapLogger *Logger = zap.NewExample()

func CreateDefaultLogWriters() WritersCollection {
	outLogWriter := Writer{IWriter: DefaultLogWriter{}}
	errorLogWriter := Writer{IWriter: DefaultErrorLogWriter{}}

	writers := WritersCollection{
		Writers:      []Writer{outLogWriter},
		ErrorWriters: []Writer{errorLogWriter},
	}

	return writers
}

func (wrs *WritersCollection) AttachWriter(w Writer) {
	wrs.Writers = append(wrs.Writers, w)
}

func (wrs *WritersCollection) AttachErrorWriter(w Writer) {
	wrs.ErrorWriters = append(wrs.ErrorWriters, w)
}
