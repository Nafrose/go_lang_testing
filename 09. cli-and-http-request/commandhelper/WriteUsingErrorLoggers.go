package commandhelper

// import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

func writeUsingErrorLoggers(
	cliBindingProperties *CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.ErrorOutputLoggers
	writUsingLoggers(cliBindingProperties, writers, line)
}
