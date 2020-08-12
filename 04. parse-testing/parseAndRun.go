package main

import (
	"flag"
	"os/exec"
	"strings"
)

func parseString(s string, separator string) string {

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
	argument := flag.String("arg", "cmd", "Write the command")
	separator := flag.String("sep", "run", "Write the separator")

	flag.Parse()

	execfile := parseString(*argument, *separator)

	cmd := exec.Command("cmd", "/c", "start", execfile)
	_ = cmd.Start()
}
