package commandhelper

// . "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
import . "github.com/nafrose/exploring/clirunner/commandhelper/core"

func writUsingLoggers(
	cliBindingProperties *CliBindingProperties,
	writers []*Writer,
	line string) {
	for _, writer := range writers {
		(*writer).Write(cliBindingProperties, line)
	}
}
