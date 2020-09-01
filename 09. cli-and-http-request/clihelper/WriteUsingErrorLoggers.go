package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/Structs"

func WriteUsingErrorLoggers(
	cliBindingProperties CliBindingProperties,
	line string) {
	writers := cliBindingProperties.WritersCollection.ErrorOutputLoggers
	wc := cliBindingProperties.WritersCollection.WriterConfiguration
	WritUsingLoggers(writers, wc, line)
}
