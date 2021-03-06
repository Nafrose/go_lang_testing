package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)
//
//type firstStepArgs struct {
//	argNumber1, argNumber2 string
//}
//
//type arguments struct {
//	runProcess                         int
//	title, msg1, msg2, delay, runTimes string
//}

func parseString(s *string) string {
	separator := "\\n"
	len := len(separator)
	stringToParse := string(*s)
	indexOfFirstSeparator := strings.Index(stringToParse, separator)
	stringAfterFirstSeparator := stringToParse[indexOfFirstSeparator+len:]

	indexOfSecondSeparator := strings.Index(stringAfterFirstSeparator, separator)
	if indexOfSecondSeparator < 0 {
		panic("Please enter arguments in correct format")
	}

	stringFromFirstToSecondSeparator := stringAfterFirstSeparator[:indexOfSecondSeparator]
	stringAfterSecondSeparator := stringAfterFirstSeparator[indexOfSecondSeparator+len:]

	var stringsAfterParsing firstStepArgs
	stringsAfterParsing.argNumber1 = stringFromFirstToSecondSeparator
	stringsAfterParsing.argNumber2 = stringAfterSecondSeparator

	var combinedStringAfterParsing string = stringsAfterParsing.argNumber1 + "," + stringsAfterParsing.argNumber2

	return combinedStringAfterParsing
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func writeToFile(s string) {
	var mutex sync.Mutex
	stringProvidedeInByte := []byte(s)

	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			err := ioutil.WriteFile("outputWrittenByCLIRunner.txt", stringProvidedeInByte, 0644)
			if err != nil {
				panic(err)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	argumentsProvided := flag.String("run", "Run,Title,Message 1,Message 2,Stream Delay,Run Times", "Write the argument")
	flag.Parse()
	// argStr := string(*argumentsProvided)
	argsProvided := parseString(argumentsProvided)

	cmd := exec.Command("cli-streamer", "-args", argsProvided)

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdoutIn and stderrIn.
	// wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		log.Println(string(stdout))
		log.Println("Trying to write it out.")
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	if errStderr != nil {
		fmt.Printf("err:\n%s\n", errStr)
	}
	writeToFile(outStr)
}
