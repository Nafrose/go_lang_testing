package commandhelper

type CmdRunningInfo struct {
	Title, Description                  string
	ID, ParentId, SessionId, BufferSize int
	IsAsync                             bool
}
