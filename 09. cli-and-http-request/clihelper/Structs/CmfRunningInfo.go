package clihelper

type CmdRunningInfo struct {
	Title, Description                  string
	ID, ParentId, SessionId, BufferSize int
	isAsync                             bool
}
