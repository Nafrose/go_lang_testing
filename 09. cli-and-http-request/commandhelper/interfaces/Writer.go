package clihelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

type Writer interface {
	Writer(cliBindingProperties *CliBindingProperties, line string)
}
