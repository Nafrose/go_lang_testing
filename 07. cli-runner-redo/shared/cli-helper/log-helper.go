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

}

func  defaultLogWriterConfig(config IWritingConfiguration) IProcessConfig {

}

func  defaultZapLogWriterConfig(config IWritingConfiguration) IProcessConfig {
	
}