package clihelper

import (
	"log"
	"sync"
)

func CliBind(c CliBindingProperties, s stdInParameter) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout := AttachLoggers(c, s)
		log.Println(string(stdout))
		log.Println("Trying to write it out.")
		wg.Done()
	}()
}
