package clihelper

type Writer interface {
	Writer(wc WriterConfiguration, line string)
}
