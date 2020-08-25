package cliHelper

import (
	_ "flag"
	"fmt"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "os/exec"
	_ "sync"
)

func CreateDefaultLogWriterConfig(processConfigProperties ProcessConfigProperties) ProcessConfig {
	if processConfigProperties.Config.IsJsonLog {
		// write for JSON
		return defaultLogJsonWriterConfig(config)
	}

	// do for non json
	return defaultLogWriterConfig(config)
}

func defaultLogJsonWriterConfig(processConfigProperties ProcessConfigProperties) ProcessConfig {
	fmt.Println("write using JSON log")

	return CreateProcessConfig(processConfigProperties, CreateDefaultLogWriters())
}

func defaultLogWriterConfig(processConfigProperties ProcessConfigProperties) ProcessConfig {
	fmt.Println("write using log")
	return CreateProcessConfig(processConfigProperties, CreateDefaultLogWriters())
}

/*
func  defaultZapLogWriterConfig(processConfigProperties ProcessConfigProperties) ProcessConfig {
	fmt.Println("write using zap")
	defaultJsonWriter := CreateDefaultWriter()
	Writers := WritersCollection{defaultJsonWriter, defaultJsonWriter}
	return CreateProcessConfig(processConfigProperties,Writers )
} */
