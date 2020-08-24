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

func CreateDefaultLogWriter(config IWritingConfiguration) {
	if config.isJsonLog {
		// write for JSON
		retun
	}

	// do for non json

}

func CreateDefaultErrorLogWriter(config IWritingConfiguration) {

}
