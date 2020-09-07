package commandloggers

import (
	"log"

	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

type DefaultOutputWriter struct{}

func (receiver DefaultOutputWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Println(line)
}
