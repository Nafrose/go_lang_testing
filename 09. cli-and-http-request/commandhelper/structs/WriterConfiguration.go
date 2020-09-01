package commandhelperstructs

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFileLocation  string
}
