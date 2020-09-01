package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
)

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFileLocation  string
}
