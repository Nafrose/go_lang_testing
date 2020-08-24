package cliHelper

type ProcessConfigProperties struct {
	title             string
	description       string
	parentProcessName string
	id                string
	parentId          string
	globalSessionId   string
	config            IWritingConfiguration
}

type IProcessConfigProperties struct {
	title             string
	description       string
	parentProcessName string
	id                string
	parentId          string
	globalSessionId   string
	config            IWritingConfiguration
}

type ProcessConfig struct {
	writers    IWritersCollection
	properties IProcessConfigProperties
}

type DefaultJsonInteration struct {
	Attributes struct {
		Counter int `json:"counter"`
		Log     struct {
			Debug     []map[string]string `json:"debug"`
			Display   []string            `json:"display"`
			Error     []map[string]string `json:"error"`
			Info      []map[string]string `json:"info"`
			IsDebug   bool                `json:"isDebug"`
			IsError   bool                `json:"isError"`
			IsInfo    bool                `json:"isInfo"`
			IsWarning bool                `json:"isWarning"`
			Warning   []map[string]string `json:"warning"`
		} `json:"log"`
		ParentTask   string `json:"parentTask"`
		ParentTaskID string `json:"parentTaskId"`
		Progress     int    `json:"progress"`
		Timestamp    string `json:"timestamp"`
	} `json:"attributes"`
	ID      string `json:"id"`
	Invoker string `json:"Invoker"`
	Title   string `json:"Title"`
}
