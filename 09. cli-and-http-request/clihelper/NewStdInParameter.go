package clihelper

func NewStdInParameters(cmd *Cmd) *StdInParameter {
	outIn, _ = cmd.StdoutPipe()
	errIn, _ = cmd.StderrPipe()
	return *StdInParameter{outIn, errIn}
}
