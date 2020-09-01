package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/builders"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

func NewWritersCollectionBuilder() *WritersCollectionBuilder {
	return &WritersCollectionBuilder{WritersCollection: &WritersCollection{}}
}

func NewCliBindingPropertiesBuilder() *CliBindingPropertiesBuilder {
	return &CliBindingPropertiesBuilder{CliBindingProperties: &CliBindingProperties{}}
}
