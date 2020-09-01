package clihelper

import "io"

type StdInParameter struct {
	StdoutIn io.ReadCloser
	StderrIn io.ReadCloser
}
