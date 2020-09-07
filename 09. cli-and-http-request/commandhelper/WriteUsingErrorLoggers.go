package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/core"

func writeUsingErrorLoggers(
	cliBindingProperties *CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.ErrorOutputLoggers
	writUsingLoggers(cliBindingProperties, writers, line)
}
