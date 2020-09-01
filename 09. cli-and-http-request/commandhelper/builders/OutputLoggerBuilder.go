package builders

import "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"

type OutputLoggerBuilder struct {
	*WritersCollectionBuilder
}

func (olb *OutputLoggerBuilder) Add(w *clihelper.Writer) *OutputLoggerBuilder {
	olb.WritersCollectionBuilder.WritersCollection.OutputLoggers =
		append(olb.WritersCollectionBuilder.WritersCollection.OutputLoggers, w)

	return olb
}

func (olb *OutputLoggerBuilder) AddMany(writers ...*clihelper.Writer) *OutputLoggerBuilder {
	for _, writer := range writers {
		olb.Add(writer)
	}

	return olb
}
