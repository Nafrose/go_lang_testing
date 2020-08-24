package cliHelper

import (
	"log"

	"go.uber.org/zap"
)

type IWriter interface {
	Write(line string)
}

type Writer struct {
	IWriter
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

func defaultWriteLog(line string) {
	log.Println(line)
}

//func defaultWriteZapLog(line string) {
//	zapLogger.info(line)
//}

func CreateDefaultLogWriter() Writer {
	return Writer{IsError: false, write: defaultWriteLog}
}

func CreateDefaultErrorLogWriter() Writer {
	return Writer{IsError: true, write: defaultWriteLog}
}

func CreateDefaultLogWriters() WritersCollection {
	outLogWriter := CreateDefaultLogWriter()
	errorLogWriter := CreateDefaultErrorLogWriter()

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

//func CreateDefaultZapWriter() Writer {
//	return Writer{IsError: false, write: defaultWriteZapLog}
//}
//
//func CreateDefaultZapErrorWriter() Writer {
//	return Writer{IsError: true, write: defaultWriteZapLog}
//}
