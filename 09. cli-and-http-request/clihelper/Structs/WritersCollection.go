package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/Interfaces"

type WritersCollection struct {
	OutputLoggers       []Writer
	ErrorOutputLoggers  []Writer
	WriterConfiguration WriterConfiguration
}
