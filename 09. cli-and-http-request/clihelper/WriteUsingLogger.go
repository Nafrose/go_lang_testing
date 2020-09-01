package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/clihelper/structs"
)

func WritUsingLoggers(
	cliBindingProperties CliBindingProperties,
	writers []Writer,
	line string) {
	for _, writer := range writers {
		writer.Writer(cliBindingProperties, line)
	}
}
