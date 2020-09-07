package commandhelper

import (
	. "os/exec"

	. "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

type CliBindingPropertiesBuilder struct {
	cliBindingProperties *CliBindingProperties
}

func (
	cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) SetWriterCollection(
	writersCollection *WritersCollection) *CliBindingPropertiesBuilder {
	cliBindingPropertiesBuilder.cliBindingProperties.WritersCollection = writersCollection

	return cliBindingPropertiesBuilder
}

func (
	cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) SetExecutor(
	cmd *Cmd) *CliBindingPropertiesBuilder {
	cliBindingPropertiesBuilder.cliBindingProperties.Cmd = cmd

	return cliBindingPropertiesBuilder
}

func (cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) Build() *CliBindingProperties {
	return cliBindingPropertiesBuilder.cliBindingProperties
}
