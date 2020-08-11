package main

import (
	"fmt"
	"strings"
)

func parse(s string, separator string) string {

	var indexOfSeparator int = strings.Index(s, separator)
	len := 1 + len(separator)
	var textObtained string = removeTextBeforeSeparator(s, indexOfSeparator, len)

	return textObtained
}

func removeTextBeforeSeparator(s string, indexOfSeparator int, lengthOfSeparator int) string {

	resultingText := s[indexOfSeparator+lengthOfSeparator:]

	return resultingText
}

func main() {

	fmt.Println(parse("command-> run cmd", "run"))

}
