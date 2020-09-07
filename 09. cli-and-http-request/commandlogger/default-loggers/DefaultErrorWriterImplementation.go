package commandloggers

import "log"

type DefaultErrorWriter struct{}

func (receiver DefaultErrorWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Fatalln(line)
}
