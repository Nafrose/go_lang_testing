package clihelper

import "log"

func RunAsync(c CliBindingProperties) StdInParameter {
	var stdout, stderr []byte
	var stdInParameter StdInParameter
	var errStdout, errStderr error

	stdInParameter := NewStdInParameters(c.Cmd)
	err := c.Cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	return stdInParameter
}
