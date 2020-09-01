package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

func CliBind(c *CliBindingProperties, s *StdInParameter) {
	go AttachLoggers(c, s)
}
