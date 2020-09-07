package commandhelper

import (
	. "os/exec"

	. "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

type CliBindingPropertiesBuilder struct {
	CliBindingProperties *CliBindingProperties
}

func (
	cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) SetWriterCollection(
	writersCollection *WritersCollection) *CliBindingPropertiesBuilder {
	cliBindingPropertiesBuilder.CliBindingProperties.WritersCollection = writersCollection

	return cliBindingPropertiesBuilder
}

func (
	cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) SetExecutor(
	cmd *Cmd) *CliBindingPropertiesBuilder {
	cliBindingPropertiesBuilder.CliBindingProperties.Cmd = cmd

	return cliBindingPropertiesBuilder
}

func (cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) Build() *CliBindingProperties {
	return cliBindingPropertiesBuilder.CliBindingProperties
}
