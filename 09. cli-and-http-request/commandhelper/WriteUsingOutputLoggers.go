package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/core"

func writeUsingOutputLoggers(
	cliBindingProperties *CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.OutputLoggers
	writUsingLoggers(cliBindingProperties, writers, line)
}
