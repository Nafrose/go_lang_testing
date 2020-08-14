package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type arguments struct {
	runProcess int
	title      string
	msg1       string
	msg2       string
	delay      int
	runTimes   int
}

func parseString(s string) (string, int) {
	var indexOfSeparator int = strings.Index(s, "\\n")
	stringAfterParsing := s[indexOfSeparator+len("\\n"):]
	// runProcessString := s[indexOfSeparator+len("\\n") : 1]
	runProcess := 2

	return stringAfterParsing, runProcess
}

func outputFromCSVFile(s string, runProcess int) arguments {
	var arguments arguments

	for i := 0; i < runProcess; i++ {
		stringAfterSplit := strings.Split(s, ",")
		delay := stringAfterSplit[4]
		streamDelay, _ := strconv.Atoi(delay)
		iteration := stringAfterSplit[5]
		runTimes, _ := strconv.Atoi(iteration)

		arguments.runProcess = runProcess
		arguments.title = stringAfterSplit[1]
		arguments.msg1 = stringAfterSplit[2]
		arguments.msg2 = stringAfterSplit[3]
		arguments.delay = streamDelay
		arguments.runTimes = runTimes
	}

	return arguments
}

func streamingMultipleTimes(arg arguments) {
	incident := 1

	for i := 0; i < arg.runTimes; i++ {
		fmt.Printf("Running process %v %s %v -> %s \n", i, arg.title, incident, arg.msg1)
		incident++
		time.Sleep(time.Duration(arg.delay) * time.Second)
		fmt.Printf("Running process %v %s %v -> %s \n", i, arg.title, incident, arg.msg2)
		incident++
	}

}

func main() {
	argumentsProvided := flag.String("args", "\nTitle,Message 1,Message 2,2,3", "Write the argument")
	flag.Parse()
	argumentString := *argumentsProvided

	argumentsStringAfterParsing, runProcess := parseString(argumentString)

	argumentsForStreaming := outputFromCSVFile(argumentsStringAfterParsing, runProcess)

	streamingMultipleTimes(argumentsForStreaming)
}
