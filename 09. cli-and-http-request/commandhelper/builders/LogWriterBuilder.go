package commandhelperbuilders

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
)

type LogWriterBuilder struct {
	*WritersCollectionBuilder
}

func (lwb *LogWriterBuilder) Add(logWriter *LogWriter) *LogWriterBuilder {
	lwb.OutputLogger().Add((*logWriter).OutputWriter())
	lwb.ErrorOutputLogger().Add((*logWriter).ErrorWriter())

	return lwb
}

func (lwb *LogWriterBuilder) AddMany(logWriters ...*LogWriter) *LogWriterBuilder {
	for _, logWriter := range logWriters {
		lwb.Add(logWriter)
	}

	return lwb
}

func (lwb *LogWriterBuilder) Build() *WritersCollection {
	return lwb.WritersCollection
}
