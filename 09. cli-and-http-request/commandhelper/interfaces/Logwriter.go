package commandhelperinterfaces

type LogWriter interface {
	OutputWriter() *Writer
	ErrorWriter() *Writer
}
