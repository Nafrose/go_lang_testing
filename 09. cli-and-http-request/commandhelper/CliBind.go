package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

func cliBind(c *CliBindingProperties, s *StdInParameter) {
	go attachCliBindingPropertiesLoggers(c, s)
}
