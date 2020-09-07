package commandloggers

import (
	"log"

	. "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

type DefaultOutputWriter struct{}

func (receiver DefaultOutputWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Println(line)
}
