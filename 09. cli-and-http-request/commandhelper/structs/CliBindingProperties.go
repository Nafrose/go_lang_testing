package clihelper

import (
	. "os/exec"
)

type CliBindingProperties struct {
	WritersCollection *WritersCollection
	CmdRunningInfo    *CmdRunningInfo
	Cmd               *Cmd
}
