package commandhelper

type LogWriter interface {
	OutputWriter() *Writer
	ErrorWriter() *Writer
}
