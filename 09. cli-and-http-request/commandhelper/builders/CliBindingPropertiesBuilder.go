package builders

import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

type CliBindingPropertiesBuilder struct {
	CliBindingProperties *CliBindingProperties
}

func NewCliBindingPropertiesBuilder() *CliBindingPropertiesBuilder {
	return &CliBindingPropertiesBuilder{&CliBindingProperties{}}
}

func (
	cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) SetWriterCollection(
	writersCollection *WritersCollection) *CliBindingPropertiesBuilder {
	cliBindingPropertiesBuilder.CliBindingProperties.WritersCollection = writersCollection

	return cliBindingPropertiesBuilder
}

func (cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) Build() *CliBindingProperties {
	return cliBindingPropertiesBuilder.CliBindingProperties
}
