package cliRunner

import cliHelper "github.com/nf/go-lang-testing/shared/cli-helper"

func RunAsync(cmd *Cmd, config cliHelper.ProcessConfig) {
	cmd.Run()
	cliHelper.BindOutputWithProcessConfigWriters(config.Writers.Writers)
}
