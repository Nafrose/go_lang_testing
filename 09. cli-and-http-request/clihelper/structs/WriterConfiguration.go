package clihelper

import "github.com/nafrose/exploring/clirunner/clihelper"

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFileLocation  string
}

type LoggerBuilder struct {
	*clihelper.CliBindingPropertiesBuilder
}

func (receiver *WriterConfiguration) OutputLoggers() *LoggerBuilder {
	// output builder return
	return &LoggerBuilder{}
}
