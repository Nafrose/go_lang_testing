package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseString(s string) string {

	var indexOfSeparator int = strings.Index(s, "\\n")
	stringAfterParsing := s[indexOfSeparator+len("\\n"):]

	return stringAfterParsing
}

func outputFromCSVFile(s string) []string {
	stringAfterSplit := strings.Split(s, ",")
	title := stringAfterSplit[0]
	msg1 := stringAfterSplit[1]
	msg2 := stringAfterSplit[2]
	delay := stringAfterSplit[3]
	iteration := stringAfterSplit[4]

	argumentTitles := []string{title, msg1, msg2, delay, iteration}

	return argumentTitles
}

func streamingMultipleTimes(title, msg1, msg2, delay, iteration string) {
	streamDelay, _ := strconv.Atoi(delay)
	runTimes, _ := strconv.Atoi(iteration)
	incident := 1

	for i := 0; i < runTimes; i++ {
		fmt.Printf("%s %v -> %s \n", title, incident, msg1)
		fmt.Printf("%s %v -> %s \n", title, incident, msg2)
		delayDurating := time.Duration(streamDelay) * time.Second
		fmt.Printf("\n%s for %d second(s)...\n\n", "Waiting...", streamDelay)

		incident++
		time.Sleep(delayDurating)
	}
}

func main() {
	argumentsProvided := flag.String("args", "Title,Message 1,Message 2,Stream Delay,Run Times", "Write the argument")
	flag.Parse()
	argumentString := *argumentsProvided
	argumentsStringAfterParsing := parseString(argumentString)
	argumentsForStreaming := outputFromCSVFile(argumentsStringAfterParsing)

	streamingMultipleTimes(
		argumentsForStreaming[0],
		argumentsForStreaming[1],
		argumentsForStreaming[2],
		argumentsForStreaming[3],
		argumentsForStreaming[4])

}
