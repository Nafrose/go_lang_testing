package cliHelper

import (
	"log"

	"go.uber.org/zap"
)

type WritersCollection struct {
	Writers      []IWriter
	ErrorWriters []IWriter
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

func defaultWriteZapLog(line string) {
	zapLogger.info(line)
}

func CreateDefaultWriter() IWriter {
	return Writer{IsError: false, write: defaultWriteLog}
}

func CreateDefaultErrorWriter() IWriter {
	return Writer{IsError: true, write: defaultWriteLog}
}

func CreateDefaultZapWriter() IWriter {
	return Writer{IsError: false, write: defaultWriteZapLog}
}

func CreateDefaultZapErrorWriter() IWriter {
	return Writer{IsError: true, write: defaultWriteZapLog}
}
