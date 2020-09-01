package commandhelper

import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

type WritersCollectionBuilder struct {
	*WritersCollection
}

func (wcb *WritersCollectionBuilder) OutputLogger() *OutputLoggerBuilder {
	return &OutputLoggerBuilder{WritersCollectionBuilder: wcb}
}

func (wcb *WritersCollectionBuilder) LogWriter() *LogWriterBuilder {
	return &LogWriterBuilder{WritersCollectionBuilder: wcb}
}

func (wcb *WritersCollectionBuilder) ErrorOutputLogger() *ErrorLoggerBuilder {
	return &ErrorLoggerBuilder{WritersCollectionBuilder: wcb}
}

func (wcb *WritersCollectionBuilder) SetConfig(writerConfig *WriterConfiguration) *WritersCollectionBuilder {
	wcb.WritersCollection.WriterConfiguration = writerConfig

	return wcb
}

func (wcb *WritersCollectionBuilder) Build() *WritersCollection {
	return wcb.WritersCollection
}
