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

func CreateDefaultErrorLogWriter(config IWritingConfiguration) {
	fmt.Println("error log writer")
}
