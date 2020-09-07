package commandhelper

// import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

type Writer interface {
	Write(cliBindingProperties *CliBindingProperties, line string)
}
