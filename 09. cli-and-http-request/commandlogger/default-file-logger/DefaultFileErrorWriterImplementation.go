package commandloggers

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
	"io/ioutil"
)

type DefaultFileErrorWriterImplementation struct{}

func (receiver *DefaultFileErrorWriterImplementation) Write(
	cliBindingProperties *CliBindingProperties,
	line string) {
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

