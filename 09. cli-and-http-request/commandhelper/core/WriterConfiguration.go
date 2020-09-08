package core

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFileLocation  string
	LogFilePath          LogFilePath
}

type LogFilePath struct {
	OutputFilePath   string
	ErrorLogFilePath string
}
