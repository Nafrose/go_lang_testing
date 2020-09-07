package commandhelper

import (
	"io"
	. "os/exec"
)

type StdInParameter struct {
	StdoutIn *io.ReadCloser
	StderrIn *io.ReadCloser
}

func NewStdInParameters(cmd *Cmd) *StdInParameter {
	outIn, _ := cmd.StdoutPipe()
	errIn, _ := cmd.StderrPipe()

	return &StdInParameter{&outIn, &errIn}
}
