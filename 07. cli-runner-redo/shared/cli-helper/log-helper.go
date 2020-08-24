package cliHelper

import (
	_ "flag"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "os/exec"
	_ "sync"
)

func CreateDefaultLogWriterConfig(processConfigProperties IProcessConfigProperties) IProcessConfig {
	if config.isJsonLog {
		// write for JSON
		retun defaultLogJsonWriterConfig(config)
	}

	// do for non json
	retun defaultLogWriterConfig(config)
}

func  defaultLogJsonWriterConfig(processConfigProperties IProcessConfigProperties) IProcessConfig {
	fmt.Println("write using JSON log")
	defaultJsonWriter := CreateDefaultWriter
	writers := 
	return CreateProcessConfig(processConfigProperties, )
}

func  defaultLogWriterConfig(processConfigProperties IProcessConfigProperties) IProcessConfig {
	fmt.Println("write using log")
	processConfigProperties = 
}

func  defaultZapLogWriterConfig(processConfigProperties IProcessConfigProperties) IProcessConfig {
	fmt.Println("write using zap")
}