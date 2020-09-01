package commandloggers

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

type DefaultFileLogWriter struct {
	*LogWriterImplementation
}

func NewDefaultFileLogWriter() *DefaultFileLogWriter {
	var outWriter Writer = &DefaultFileOutputWriterImplementation{}
	var errorWriter Writer = &DefaultFileErrorWriterImplementation{}

	return &DefaultFileLogWriter{LogWriterImplementation: NewLogWriter(&outWriter, &errorWriter)}
}

