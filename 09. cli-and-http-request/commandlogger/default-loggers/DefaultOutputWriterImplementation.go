package commandloggers

import "log"

type DefaultOutputWriter struct{}

func (receiver DefaultOutputWriter) Write(cliBindingProperties *CliBindingProperties, line string) {
	log.Println(line)
}
