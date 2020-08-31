package clihelper

import "log"

func RunAsync(c CliBindingProperties) stdInParameter {
	var stdout, stderr []byte
	var stdInParameter stdInParameter
	var errStdout, errStderr error
	stdInParameter.stdoutIn, _ = c.Cmd.StdoutPipe()
	stdInParameter.stderrIn, _ = c.Cmd.StderrPipe()
	err := c.Cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	return stdInParameter
}
