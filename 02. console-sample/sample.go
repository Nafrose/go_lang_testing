package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// How to take line from console : https://bit.ly/3iqK0nc, https://bit.ly/2DMnpm8
func GetLine(promptMessage string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if !IsEmpty(promptMessage) {
		fmt.Print(promptMessage)
	}

	// input, er := reader.ReadString('\n')
	// input2 := strings.TrimRight(input, "\n")

	return reader.ReadString('\n')
}
