package clihelper

type CmdRunner struct {
	CliBindingProperties CliBindingProperties
}

type Writer interface {
	Writer(wc WriterConfiguration, line string)
}

type StdInParameter struct {
	stdoutIn []byte
	stderrIn error
}

type Add interface {
	Add(array []Writer)
}
