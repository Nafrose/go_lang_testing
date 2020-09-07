package main

import (
	"os/exec"

	cmdhelper "github.com/nafrose/exploring/clirunner/commandhelper"
	commandhelper "github.com/nafrose/exploring/clirunner/commandhelper/core"
	commandloggers "github.com/nafrose/exploring/clirunner/commandlogger/default-loggers"
)

const (
	CliStreamerExec = "cli-streamer.exe"
)

func runCli() {
	writersCollection := NewWritersCollection()
	cliBindingProperties := GetCliBindingProperties(writersCollection)
	cmdhelper.RunAsync(cliBindingProperties)
}

func GetCliBindingProperties(
	writersCollection *commandhelper.WritersCollection) *commandhelper.CliBindingProperties {
	cmd := exec.Cmd{
		Path: CliStreamerExec,
		Args: nil,
	}

	cliBindingProperties := cmdhelper.
		NewCliBindingPropertiesBuilder().
		SetWriterCollection(writersCollection).
		SetExecutor(&cmd).
		Build()

	return cliBindingProperties
}

func NewWritersCollection() *commandhelper.WritersCollection {
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
