package commandhelper

// . "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
// . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

func writUsingLoggers(
	cliBindingProperties *CliBindingProperties,
	writers []*Writer,
	line string) {
	for _, writer := range writers {
		(*writer).Write(cliBindingProperties, line)
	}
}
