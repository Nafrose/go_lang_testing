package commandhelper

import commandhelper "github.com/nafrose/exploring/clirunner/commandhelper/core"

// import "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"

type ErrorLoggerBuilder struct {
	*WritersCollectionBuilder
}

func (elb *ErrorLoggerBuilder) Add(w *commandhelper.Writer) *ErrorLoggerBuilder {
	elb.WritersCollectionBuilder.WritersCollection.ErrorOutputLoggers =
		append(elb.WritersCollectionBuilder.WritersCollection.ErrorOutputLoggers, w)

	return elb
}

func (elb *ErrorLoggerBuilder) AddMany(writers ...*commandhelper.Writer) *ErrorLoggerBuilder {
	for _, writer := range writers {
		elb.Add(writer)
	}

	return elb
}
