package commandloggers

import (
	"log"

	. "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

type DefaultErrorWriter struct{}

func (receiver DefaultErrorWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Fatalln(line)
}
