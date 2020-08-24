package cliHelper

type IWriter interface{
	IsError bool
	Write(line string)
}

type IWritersCollection interface{
	Writers []IWriter
	ErrorWriters []IWriter
}

type IWritingConfiguration interface{
	IsJsonLog	bool
	// if enabled then write to fmt.Fprintf
	IsWriteToHttp bool
	// https://github.com/uber-go/zap
	IsLogUsingZap bool
	IsWriteToFile bool
	WriteToFileDirectory string
	WriteToFilePath string
	IsWriteToDirectory bool
}