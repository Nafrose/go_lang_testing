package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

func WriteUsingOutputLoggers(
	cliBindingProperties CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.OutputLoggers
	WritUsingLoggers(cliBindingProperties, writers, line)
}
