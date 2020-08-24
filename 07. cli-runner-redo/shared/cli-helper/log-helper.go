package cliHelper

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

func CreateDefaultLogWriter(config IWritingConfiguration) {
	if config.isJsonLog {
		
	}
}

func CreateDefaultErrorLogWriter(config IWritingConfiguration) {

}
