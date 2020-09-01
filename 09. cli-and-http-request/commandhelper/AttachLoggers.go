package commandhelper

import (
	. "github.com/nafrose/exploring/clirunner/commandhelper/structs"
	"io"
)

type LogWrite func(cliBindingProperties *CliBindingProperties, line string)

func AttachLoggers(
	cliBindingProperties *CliBindingProperties,
	stdInParameter *StdInParameter) {
	if cliBindingProperties.CmdRunningInfo.IsAsync {
		go attachInternalLoggers(cliBindingProperties, stdInParameter.StdoutIn, WriteUsingOutputLoggers)
		go attachInternalLoggers(cliBindingProperties, stdInParameter.StderrIn, WriteUsingErrorLoggers)

		return
	}

	go attachInternalLoggers(cliBindingProperties, stdInParameter.StderrIn, WriteUsingErrorLoggers)
	attachInternalLoggers(cliBindingProperties, stdInParameter.StdoutIn, WriteUsingOutputLoggers)
}

func attachInternalLoggers(
	cliBindingProperties *CliBindingProperties,
	readCloser *io.ReadCloser,
	logWrite LogWrite) ([]byte, error) {
	var out []byte
	buff := make([]byte, 1, cliBindingProperties.CmdRunningInfo.BufferSize)
	n, err := (*readCloser).Read(buff[:])
	if n > 0 {
		bufferedSplits := buff[:n]
		out = append(out, bufferedSplits...)
		line := string(bufferedSplits)
		logWrite(cliBindingProperties, line)
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
