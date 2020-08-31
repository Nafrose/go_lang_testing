package clihelper

import (
	"log"
)

func CliBind(c CliBindingProperties, s StdInParameter) {
		stdout, errStdout := go AttachLoggers(c, s)
		log.Println(string(stdout))
		log.Println("Trying to write it out.")

}
