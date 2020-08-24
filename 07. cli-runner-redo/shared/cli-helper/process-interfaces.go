package cliHelper

type IProcessConfig interface {
	writers           IWritersCollection
	config            IWritingConfiguration
	title             string
	description       string
	parentProcessName string
}

type  IDefaultJsonInteration interface{
	Attributes	struct {
		Counter	int
		Log	struct {
			Debug		[]map[string]string
			Display		[]string
			Error		[]map[string]string
			Info		[]map[string]string
			IsDebug		bool
			IsError		bool
			IsInfo		bool
			IsWarning	bool
			Warning		[]map[string]string
		}
		ParentTask	string
		ParentTaskID	string
		Progress	int
		Timestamp	string
	}
	ID	string
	Invoker	string
	Title	string
}
