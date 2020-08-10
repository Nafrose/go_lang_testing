package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// How to take line from console : https://bit.ly/3iqK0nc, https://bit.ly/2DMnpm8
func getLine(promptMessage string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if !isEmpty(promptMessage) {
		fmt.Print(promptMessage)
	}

	// input, er := reader.ReadString('\n')
	// input2 := strings.TrimRight(input, "\n")

	return reader.ReadString('\n')
}

func main() {
	s, _ := getLine("Please enter your text to revserse:")
	fmt.Print("entered : ")
	fmt.Print(s)
	var r1, r2 = reverse(s)
	fmt.Printf("input: %s\n reverse1: %s \n reverse2: %s \n", s, r1, r2)
}

func reverse(s string) (string, string) {
	len := len(s)
	var r = make([]byte, len)
	var r2 = make([]byte, len)
	mid := int(len / 2)

	for i := 0; i < len; i++ {
		r[i] = s[len-1-i]
		fmt.Printf("%c-%c\n", s[i], s[len-1-i])
	}

	for i2 := 0; i2 < mid; i2++ {
		lastIndex := len - 1 - i2
		r2[i2] = s[lastIndex]
		r2[lastIndex] = s[i2]
	}

	r2[mid] = s[mid]

	reverse := string(r)
	reverse2 := string(r2)

	return reverse, reverse2
}
