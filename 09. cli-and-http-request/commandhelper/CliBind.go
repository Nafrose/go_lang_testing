package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/core"

func cliBind(c *CliBindingProperties, s *StdInParameter) {
	go attachCliBindingPropertiesLoggers(c, s)
}
