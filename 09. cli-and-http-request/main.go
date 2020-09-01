package main

import (
	cmdhelper "github.com/nafrose/exploring/clirunner/commandhelper"
	commandloggers "github.com/nafrose/exploring/clirunner/commandlogger/default-loggers"
)

func runCli() {
	writersCollection := NewWritersCollection()

	cliBindingProperties := cmdhelper.
		NewCliBindingPropertiesBuilder().
		SetWriterCollection(writersCollection).
		SetExecutor()
	Build()

	cmdhelper.RunAsync(cliBindingProperties)
}

func NewWritersCollection() *WritersCollection {
	defaultLogger := commandloggers.NewDefaultLogWriter()
	writersCollection := cmdhelper.NewWritersCollectionBuilder().
		LogWriter().
		Add(defaultLogger).
		Build()

	return writersCollection
}

func main() {
	//writersCollection := cmdhelper.WritersCollection.
	//	OutputLogger().Add(...). -> OutputLoggerBuilder
	//ErrorOutputLogger().Add(...).
	//SetConfig(WriterConfiguration{IsJson...}).
	//	Build()
	runCli()
}
