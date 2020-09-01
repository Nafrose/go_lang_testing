package commandhelperstructs

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
)

type WritersCollection struct {
	OutputLoggers       []*Writer
	ErrorOutputLoggers  []*Writer
	WriterConfiguration *WriterConfiguration
}
