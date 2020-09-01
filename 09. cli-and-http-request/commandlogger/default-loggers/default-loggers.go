package commandloggers

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
	"log"
)

type DefaultLogWriter struct{}
type DefaultErrorWriter struct{}
type DefaultOutputWriter struct{}

func (receiver *DefaultOutputWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Println(line)
}

func (receiver *DefaultErrorWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Fatalln(line)
}

func (receiver *DefaultLogWriter) OutputWriter() Writer {
	return DefaultOutputWriter{}
}

func (receiver *DefaultLogWriter) ErrorWriter() {

}
