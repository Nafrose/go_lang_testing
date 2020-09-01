package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/Structs"

type Writer interface {
	Writer(cliBindingProperties CliBindingProperties, line string)
}
