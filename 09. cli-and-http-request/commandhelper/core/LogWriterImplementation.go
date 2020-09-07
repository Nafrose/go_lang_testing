package commandhelper

// import . "github.com/nafrose/exploring/clirunner/commandhelper/interfaces"

type LogWriterImplementation struct {
	outputLogWriter *Writer
	errorLogWriter  *Writer
}

func NewLogWriter(
	outputWriter *Writer,
	errorWriter *Writer) *LogWriterImplementation {
	return &LogWriterImplementation{
		outputLogWriter: outputWriter,
		errorLogWriter:  errorWriter,
	}
}

func (logWriterImplementation *LogWriterImplementation) OutputWriter() *Writer {
	return logWriterImplementation.outputLogWriter
}

func (logWriterImplementation *LogWriterImplementation) ErrorWriter() *Writer {
	return logWriterImplementation.errorLogWriter
}
