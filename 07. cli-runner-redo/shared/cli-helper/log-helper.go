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

func CreateDefaultLogWriterConfig(config IWritingConfiguration) IProcessConfig {
	if config.isJsonLog {
		// write for JSON
		retun defaultLogJsonWriterConfig(config)
	}

	// do for non json
	retun defaultLogWriterConfig(config)
}

func  defaultLogJsonWriterConfig(config IWritingConfiguration) IProcessConfig {
	fmt.Println("write using JSON log")

	return 
}

func  defaultLogWriterConfig(config IWritingConfiguration) IProcessConfig {
	fmt.Println("write using log")

}

func  defaultZapLogWriterConfig(config IWritingConfiguration) IProcessConfig {
	fmt.Println("write using zap")
}