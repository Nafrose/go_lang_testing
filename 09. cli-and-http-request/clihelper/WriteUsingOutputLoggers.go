package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/Structs"

func WriteUsingOutputLoggers(cliBindingProperties CliBindingProperties, line string) {
	writers := cliBindingProperties.WritersCollection.OutputLoggers
	wc := cliBindingProperties.WritersCollection.WriterConfiguration
	WritUsingLoggers(writers, wc, line)
}
