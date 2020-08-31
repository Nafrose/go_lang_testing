package clihelper

func NewStdInParameters(cmd *Cmd) *StdInParameter {
	StdInParameter.stdoutIn, _ = cmd.StdoutPipe()
	StdInParameter.stderrIn, _ = cmd.StderrPipe()
	return *StdInParameter
}
