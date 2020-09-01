package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/structs"

func WriteUsingOutputLoggers(
	cliBindingProperties CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.OutputLoggers
	WritUsingLoggers(cliBindingProperties, writers, line)
}
