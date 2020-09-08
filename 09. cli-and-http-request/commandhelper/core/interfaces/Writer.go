package interfaces

import . "github.com/nafrose/exploring/clirunner/commandhelper/core"

type Writer interface {
	Write(cliBindingProperties *CliBindingProperties, line string)
}
