package clihelper

import (
	. "github.com/nafrose/exploring/clirunner/clihelper/Structs"
	"io"
)

func AttachLoggers(
	cliBindingProperties CliBindingProperties,
	stdInParameter StdInParameter) ([]byte, error) {
	attachInternalLoggers(
		cliBindingProperties,
		stdInParameter.StdoutIn,
	)
}

func attachInternalLoggers(
	cliBindingProperties CliBindingProperties,
	readCloser io.ReadCloser,
	logger WriteUsingLogger) ([]byte, error) {
	var out []byte
	buff := make([]byte, 1, cliBindingProperties.CmdRunningInfo.BufferSize)
	n, err := readCloser.Read(buff[:])
	if n > 0 {
		bufferedSplits := buff[:n]
		out = append(out, bufferedSplits...)
		line := string(bufferedSplits)
		logger.WriteUsingTheLogger(cliBindingProperties, line)
	}
	if err != nil {
		// Read returns io.EOF at the end of file, which is not an error for us
		if err == io.EOF {
			err = nil
		}
		return out, err
	}

	return nil, nil
}
