package main

import (
	"fmt"
	"os/exec"

	cmdhelper "github.com/nafrose/exploring/clirunner/commandhelper"
	commandhelper "github.com/nafrose/exploring/clirunner/commandhelper/core"
	commandloggers "github.com/nafrose/exploring/clirunner/commandlogger/default-loggers"
)

const (
	CliStreamerExec = "cli-streamer.exe"
)

var args = make([]string, 0, 10)

func runCli() {
	writersCollection := NewWritersCollection()
	cliBindingProperties := GetCliBindingProperties(writersCollection)
	cmdhelper.RunAsync(cliBindingProperties)
}

func GetCliBindingProperties(
	writersCollection *commandhelper.WritersCollection) *commandhelper.CliBindingProperties {
	cmd := exec.Cmd{
		Path: CliStreamerExec,
		Args: args,
	}
	fmt.Println("inside GetClibindingProp")
	cliBindingProperties := cmdhelper.
		NewCliBindingPropertiesBuilder().
		SetWriterCollection(writersCollection).
		SetExecutor(&cmd).
		Build()
	fmt.Println(cliBindingProperties)
	return cliBindingProperties
}

func NewWritersCollection() *commandhelper.WritersCollection {
	defaultLogger := commandloggers.NewDefaultLogWriter()
	writersCollection := cmdhelper.NewWritersCollectionBuilder().
		LogWriter().
		Add(defaultLogger).
		Build()
	fmt.Println("inside NewWriterCollection", writersCollection)
	return writersCollection
}

func main() {
	args = append(args, "2", "Title", "Message 1", "Message 2", "2", "3")
	wC := cmdhelper.NewWritersCollectionBuilder()
	wC.SetConfig().
		WriterConfiguration{ShouldWriteToFile: true, WriteToFileLocation: "./"} 
		// LogFilePath: {OutputFilePath: "./", ErrorLogFilePath: "./"}

		instance :=wC.Build()

	// wC := NewWritersCollection()
	//writersCollection := cmdhelper.WritersCollection.
	//	OutputLogger().Add(...). -> OutputLoggerBuilder
	//ErrorOutputLogger().Add(...).
	//SetConfig(WriterConfiguration{IsJson...}).
	//	Build()
	runCli()
}
