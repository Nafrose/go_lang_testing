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

func CreateDefaultErrorLogWriterConfig(processConfigProperties IProcessConfigProperties) IProcessConfig {
	fmt.Println("error log writer")
	return ProcessConfig{config: config, title}
}
