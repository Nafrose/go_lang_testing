package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type arguments struct {
	title    string
	msg1     string
	msg2     string
	delay    int
	runTimes int
}

func parseString(s string) string {
	var indexOfSeparator int = strings.Index(s, "\\n")
	stringAfterParsing := s[indexOfSeparator+len("\\n"):]

	return stringAfterParsing
}

func outputFromCSVFile(s string) arguments {
	var arguments arguments

	stringAfterSplit := strings.Split(s, ",")
	delay := stringAfterSplit[3]
	streamDelay, _ := strconv.Atoi(delay)
	iteration := stringAfterSplit[4]
	runTimes, _ := strconv.Atoi(iteration)

	arguments.title = stringAfterSplit[0]
	arguments.msg1 = stringAfterSplit[1]
	arguments.msg2 = stringAfterSplit[2]
	arguments.delay = streamDelay
	arguments.runTimes = runTimes

	return arguments
}

func streamingMultipleTimes(arg arguments) {
	incident := 1

	for i := 0; i < arg.runTimes; i++ {
		fmt.Printf("%s %v -> %s \n", arg.title, incident, arg.msg1)
		incident++
		time.Sleep(time.Duration(arg.delay) * time.Second)
		fmt.Printf("%s %v -> %s \n", arg.title, incident, arg.msg2)
		incident++
	}
}

func main() {
	var sampleArgument string = "Title,Message 1,Message 2,Stream Delay,Run Times"

	argumentsProvided := flag.String("args", sampleArgument, "Write the argument")
	flag.Parse()

	argumentString := *argumentsProvided

	argumentsStringAfterParsing := parseString(argumentString)

	argumentsForStreaming := outputFromCSVFile(argumentsStringAfterParsing)

	streamingMultipleTimes(argumentsForStreaming)

}
