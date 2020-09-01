package commandhelper

import (
	. "os/exec"
)

type CliBindingProperties struct {
	WritersCollection *WritersCollection
	CmdRunningInfo    *CmdRunningInfo
	Cmd               *Cmd
}
