package clihelper

type WritersCollection struct {
	OutputLoggers       []Writer
	ErrorOutputLoggers  []Writer
	WriterConfiguration WriterConfiguration
}
