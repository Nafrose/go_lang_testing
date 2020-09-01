package builders

import . "github.com/nafrose/exploring/clirunner/clihelper/structs"

type WritersCollectionBuilder struct {
	WritersCollection *WritersCollection
}

func NewWritersCollectionBuilder() *WritersCollectionBuilder {
	return &WritersCollectionBuilder{&WritersCollection{}}
}

func (wcb *WritersCollectionBuilder) OutputLogger() *OutputLoggerBuilder {
	return &OutputLoggerBuilder{WritersCollectionBuilder: wcb}
}
