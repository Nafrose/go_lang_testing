package commandloggers

import (
	"log"

	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

type DefaultErrorWriter struct{}

func (receiver DefaultErrorWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Fatalln(line)
}
