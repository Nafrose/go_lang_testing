package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type firstStepArgs struct {
	argNumber1, argNumber2 string
}

type arguments struct {
	runProcess, delay, runTimes int
	title, msg1, msg2           string
}

func outputFromCSVFile(s string) []arguments {
	var arguments = make([]arguments, 2)
	numberOfArgsProvided := 6

	for i := 0; i < 2; i++ {
		stringAfterSplit := strings.Split(s, ",")

		arguments[i].runProcess, _ = strconv.Atoi(stringAfterSplit[0+i*numberOfArgsProvided])
		arguments[i].title = stringAfterSplit[1+i*numberOfArgsProvided]
		arguments[i].msg1 = stringAfterSplit[2+i*numberOfArgsProvided]
		arguments[i].msg2 = stringAfterSplit[3+i*numberOfArgsProvided]
		arguments[i].delay, _ = strconv.Atoi(stringAfterSplit[4+i*numberOfArgsProvided])
		arguments[i].runTimes, _ = strconv.Atoi(stringAfterSplit[5+i*numberOfArgsProvided])
	}

	return arguments
}

func streamingMultipleTimes(arg []arguments) {
	for j := 0; j < 2; j++ {
		for i := 0; i < arg[j].runTimes; i++ {
			fmt.Printf("Running process %v %s %v -> %s \n", j+1, arg[j].title, i+1, arg[j].msg1)
			fmt.Printf("Running process %v %s %v -> %s \n", j+1, arg[j].title, i+1, arg[j].msg2)

			time.Sleep(time.Duration(arg[j].delay) * time.Second)
		}

		fmt.Printf("\n")
	}
}

func main() {
	argumentsProvided := flag.String("args", "\nrunProcess,Title,Message 1,Message 2,2,3", "Write the argument")
	flag.Parse()

	strArgumentsPRovided := *argumentsProvided
	argumentsForStreaming := outputFromCSVFile(strArgumentsPRovided)

	streamingMultipleTimes(argumentsForStreaming)
}
