package main

import (
	_ "flag"
	_ "fmt"
	cliHelper "github.com/nf/go-lang-testing/shared/cli-helper"
	_ "io"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "os/exec"
	_ "sync"
)

func main() {
	writingConfig := cliHelper.ProcessConfigProperties{
		Title:             "",
		Description:       "",
		ParentProcessName: "",
		Id:                "",
		ParentId:          "",
		GlobalSessionId:   "",
		Config:            cliHelper.WritingConfiguration{},
	}

	processConfig := cliHelper.CreateDefaultLogWriterConfig(writingConfig)
	processConfig.Writers.AttachErrorWriter()

}
