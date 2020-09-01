package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

func WritUsingLoggers(
	cliBindingProperties CliBindingProperties,
	writers []Writer,
	line string) {
	for _, writer := range writers {
		writer.Writer(cliBindingProperties, line)
	}
}
