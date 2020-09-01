package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/Structs"
)

func CliBind(c CliBindingProperties, s StdInParameter) {
	go AttachLoggers(c, s)
}
