package main

import cmdhelper "github.com/nafrose/exploring/clirunner/commandhelper"

func main() {
	//writersCollection := cmdhelper.WritersCollection.
	//	OutputLogger().Add(...). -> OutputLoggerBuilder
	//ErrorOutputLogger().Add(...).
	//SetConfig(WriterConfiguration{IsJson...}).
	//	Build()

	cmdhelper.NewWritersCollectionBuilder().OutputLogger().Add()
}
