package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/Interfaces"
	. "github.com/nafrose/exploring/clirunner/clihelper/Structs"
)

func WritUsingLoggers(
	writers []Writer,
	wc WriterConfiguration,
	line string) {
	for _, writer := range writers {
		writer.Writer(wc, line)
	}
}
