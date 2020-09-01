package clihelper

import (
	"github.com/nafrose/exploring/clirunner/clihelper/builders"
	. "github.com/nafrose/exploring/clirunner/clihelper/interfaces"
)

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFileLocation  string
}

type OutputLoggerBuilder struct {
	*builders.WritersCollectionBuilder
}

func (receiver *OutputLoggerBuilder) Add(w *Writer) *OutputLoggerBuilder {
	receiver.WritersCollectionBuilder.WritersCollection.OutputLoggers =
		append(receiver.WritersCollectionBuilder.WritersCollection.OutputLoggers, *w)

	return receiver
}
