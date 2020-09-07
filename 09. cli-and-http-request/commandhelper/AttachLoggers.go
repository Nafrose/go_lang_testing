package commandhelper

import (
	"io"

	. "github.com/nafrose/exploring/clirunner/commandhelper/core"
)

type logWrite func(cliBindingProperties *CliBindingProperties, line string)

func attachCliBindingPropertiesLoggers(
	cliBindingProperties *CliBindingProperties,
	stdInParameter *StdInParameter) {
	if cliBindingProperties.CmdRunningInfo.IsAsync {
		go attachInternalLoggers(cliBindingProperties, stdInParameter.StdoutIn, writeUsingOutputLoggers)
		go attachInternalLoggers(cliBindingProperties, stdInParameter.StderrIn, writeUsingErrorLoggers)

		return
	}

	go attachInternalLoggers(cliBindingProperties, stdInParameter.StderrIn, writeUsingErrorLoggers)
	attachInternalLoggers(cliBindingProperties, stdInParameter.StdoutIn, writeUsingOutputLoggers)
}

func attachInternalLoggers(
	cliBindingProperties *CliBindingProperties,
	readCloser *io.ReadCloser,
	logWrite logWrite) ([]byte, error) {
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
