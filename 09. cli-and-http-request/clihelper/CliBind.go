package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/structs"
)

func CliBind(c CliBindingProperties, s StdInParameter) {
	go AttachLoggers(c, s)
}
