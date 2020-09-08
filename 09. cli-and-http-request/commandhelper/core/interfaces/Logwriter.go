package interfaces

type LogWriter interface {
	OutputWriter() *Writer
	ErrorWriter() *Writer
}
