package commandloggers

import (
	"io/ioutil"
	"sync"

	commandhelper "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

type DefaultFileOutputWriterImplementation struct{}

var outputMutex = sync.Mutex{}

func (receiver *DefaultFileOutputWriterImplementation) Write(cliBindingProperties *commandhelper.CliBindingProperties, line string) {
	if !cliBindingProperties.WritersCollection.WriterConfiguration.ShouldWriteToFile {
		return
	}

	logFilePath := cliBindingProperties.WritersCollection.WriterConfiguration.LogFilePath
	stringsInBytes := []byte(line)

	outputMutex.Lock()
	{
		err := ioutil.WriteFile(logFilePath.OutputFilePath, stringsInBytes, 0644)
		if err != nil {
			panic(err)
		}
	}

	outputMutex.Unlock()
}
