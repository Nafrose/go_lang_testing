package commandhelperstructs

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
)

type WritersCollection struct {
	LogWriters          []*LogWriter
	WriterConfiguration *WriterConfiguration
}
