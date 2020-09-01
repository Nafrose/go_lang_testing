package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/builders"
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	implementations "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

func NewWritersCollectionBuilder() *WritersCollectionBuilder {
	return &WritersCollectionBuilder{WritersCollection: &implementations.WritersCollection{}}
}

func NewCliBindingPropertiesBuilder() *CliBindingPropertiesBuilder {
	return &CliBindingPropertiesBuilder{CliBindingProperties: &implementations.CliBindingProperties{}}
}

func NewLogWriter(
	outputWriter *Writer,
	errorWriter *Writer) *implementations.LogWriterImplementation {
	return implementations.NewLogWriter(outputWriter, errorWriter)
}
