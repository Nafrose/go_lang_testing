package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/Interfaces"
	. "github.com/nafrose/exploring/clirunner/clihelper/Structs"
)

func WritUsingLoggers(
	cliBindingProperties CliBindingProperties,
	writers []Writer,
	line string) {
	for _, writer := range writers {
		writer.Writer(cliBindingProperties, line)
	}
}
