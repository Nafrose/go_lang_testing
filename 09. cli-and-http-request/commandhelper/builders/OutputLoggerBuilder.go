package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/core"

type OutputLoggerBuilder struct {
	*WritersCollectionBuilder
}

func (olb *OutputLoggerBuilder) Add(w *Writer) *OutputLoggerBuilder {
	olb.WritersCollectionBuilder.WritersCollection.OutputLoggers =
		append(olb.WritersCollectionBuilder.WritersCollection.OutputLoggers, w)

	return olb
}

func (olb *OutputLoggerBuilder) AddMany(writers ...*Writer) *OutputLoggerBuilder {
	for _, writer := range writers {
		olb.Add(writer)
	}

	return olb
}
