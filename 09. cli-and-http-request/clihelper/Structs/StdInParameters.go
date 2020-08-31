package clihelper

type CmdRunner struct {
	CliBindingProperties CliBindingProperties
}

type StdInParameter struct {
	stdoutIn []byte
	stderrIn error
}
