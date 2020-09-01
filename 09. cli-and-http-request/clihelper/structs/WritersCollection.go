package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/interfaces"
)

type WritersCollection struct {
	OutputLoggers       []Writer
	ErrorOutputLoggers  []Writer
	WriterConfiguration *WriterConfiguration
}
