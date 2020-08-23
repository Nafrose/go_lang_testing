package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type firstStepArgs struct {
	argNumber1, argNumber2 string
}

type arguments struct {
	runProcess                         int
	title, msg1, msg2, delay, runTimes string
}

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

func outputFromCSVFile(s string) []arguments {
	var arguments = make([]arguments, 2)
	numberOfItemsInArguments := 6

	stringAfterSplit := strings.Split(s, ",")
	for i := 0; i < 2; i++ {
		arguments[i].runProcess, _ = strconv.Atoi(stringAfterSplit[i*numberOfItemsInArguments+0])
		arguments[i].title = stringAfterSplit[i*numberOfItemsInArguments+1]
		arguments[i].msg1 = stringAfterSplit[(i*numberOfItemsInArguments + 2)]
		arguments[i].msg2 = stringAfterSplit[i*numberOfItemsInArguments+3]
		arguments[i].delay = stringAfterSplit[i*numberOfItemsInArguments+4]
		arguments[i].runTimes = stringAfterSplit[i*numberOfItemsInArguments+5]
	}

	return arguments
}

func compileToCLIArgs(arg []arguments) string {
	var stringArg = make([]string, 2)
	for i := 0; i < 2; i++ {
		stringArg[i] = arg[i].title + "," + arg[i].msg1 + "," + arg[i].msg2 + "," + arg[i].delay + "," + arg[i].runTimes
	}

	stringAfterCompile := strings.Join(stringArg, ",")
	return stringAfterCompile
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
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
	argStr := string(*argumentsProvided)
	fmt.Println(argStr)

	argsProvided := parseString(argumentsProvided)
	// fmt.Println("argsProvided:")
	// fmt.Println(argsProvided)
	// stringArgs := outputFromCSVFile(argsProvided)
	// fmt.Println("stringArgs:")
	// fmt.Println(stringArgs)
	// argString := compileToCLIArgs(stringArgs)
	// fmt.Printf("argString: %s \n", argString)

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
