package commandloggers

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
	"log"
)

type DefaultLogWriter struct {
	*LogWriterImplementation
}

func NewDefaultLogWriter() *DefaultLogWriter {
	var outWriter Writer = DefaultOutputWriter{}
	var errorWriter Writer = DefaultErrorWriter{}

	return &DefaultLogWriter{LogWriterImplementation: NewLogWriter(&outWriter, &errorWriter)}
}

type DefaultErrorWriter struct{}
type DefaultOutputWriter struct{}

func (receiver DefaultOutputWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Println(line)
}

func (receiver DefaultErrorWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Fatalln(line)
}
