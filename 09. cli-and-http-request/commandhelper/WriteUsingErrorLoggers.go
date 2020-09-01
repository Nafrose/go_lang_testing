package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

func WriteUsingErrorLoggers(
	cliBindingProperties *CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.ErrorOutputLoggers
	WritUsingLoggers(cliBindingProperties, writers, line)
}
