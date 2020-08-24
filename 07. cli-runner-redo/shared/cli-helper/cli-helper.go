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

func CreateProcessConfig(
	properties IProcessConfigProperties,
	writers IWritersCollection) IProcessConfig {
	outputWriters, errorWriters := getOutputErrorWriters(writers)

	return ProcessConfig{outputWriters, errorWriters, properties}
}
