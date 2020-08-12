package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {

	var outputFromFile string = "Hello from cli1"

	cmd := exec.Command("cli2.exe")
	outputFromOtherFile, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	var combinedOutputFromFiles []byte = []byte(outputFromFile + " " + string(outputFromOtherFile))

	err = ioutil.WriteFile("CombinedOutputFromFiles", combinedOutputFromFiles, 0644)
	if err != nil {
		panic(err)
	}
}
