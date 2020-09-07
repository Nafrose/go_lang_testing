package commandloggers

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

type DefaultLogWriter struct {
	*LogWriterImplementation
}

func NewDefaultLogWriter() *LogWriter {
	var outWriter Writer = DefaultOutputWriter{}
	var errorWriter Writer = DefaultErrorWriter{}
	var logWriter LogWriter = &DefaultLogWriter{
		LogWriterImplementation: NewLogWriter(&outWriter, &errorWriter),
	}

	return &logWriter
}
