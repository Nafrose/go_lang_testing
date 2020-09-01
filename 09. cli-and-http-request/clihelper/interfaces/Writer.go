package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/structs"

type Writer interface {
	Writer(cliBindingProperties CliBindingProperties, line string)
}
