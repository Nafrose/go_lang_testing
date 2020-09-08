package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/builders"
	implementations "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

func NewWritersCollectionBuilder() *WritersCollectionBuilder {
	return &WritersCollectionBuilder{WritersCollection: &implementations.WritersCollection{}}
}

func NewCliBindingPropertiesBuilder() *CliBindingPropertiesBuilder {
	return &CliBindingPropertiesBuilder{CliBindingProperties: &implementations.CliBindingProperties{}}
}

func CreateLogWriter(
	outputWriter *implementations.Writer,
	errorWriter *implementations.Writer) *implementations.LogWriterImplementation {
	return implementations.NewLogWriter(outputWriter, errorWriter)
}
