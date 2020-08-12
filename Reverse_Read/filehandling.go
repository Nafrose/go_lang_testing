package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const lineSeparator = "\n"

func reverseStringLine(s string) string {
	length := len(s)
	var r2 = make([]byte, length)
	mid := length / 2

	for i2 := 0; i2 < mid; i2++ {
		lastIndex := length - 1 - i2
		r2[i2] = s[lastIndex]
		r2[lastIndex] = s[i2]
	}

	if length%2 == 1 {
		r2[mid] = s[mid]
	}

	reverse2 := string(r2)

	return reverse2
}

func fileTextReverseByLine(s []byte) string {
	lenOf := len(s)
	str := string(s)

	lines := strings.Split(str, lineSeparator)
	compiledLines := make([]string, 0, lenOf)

	for _, line := range lines {
		reverseLine := strings.TrimSpace(reverseStringLine(line))
		compiledLines = append(compiledLines, reverseLine)
	}

	return strings.Join(compiledLines, lineSeparator)
}

func readCurrentDir() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}

	defer file.Close()

	fileList, _ := file.Readdir(0)

	fmt.Printf("\nName\t\tSize\tIsDirectory  Last Modification\n")

	for _, files := range fileList {
		fmt.Printf("\n%-15s %-7v %-12v %v", files.Name(), files.Size(), files.IsDir(), files.ModTime())
	}
}

func main() {
	// read the whole file at once
	readCurrentDir()

	currentDirectory, _ := os.Getwd() //filepath.Abs(filepath.Dir(os.Args[0]))
	processingFilePath := filepath.Join(currentDirectory, "test.txt")
	fmt.Printf("\nReading file: %s\n", processingFilePath)

	b, err := ioutil.ReadFile(processingFilePath)

	if err != nil {
		panic(err)
	}

	reversedTextByLines := []byte(fileTextReverseByLine(b))

	//write the whole body at once
	err = ioutil.WriteFile("output.txt", reversedTextByLines, 0644)
	if err != nil {
		panic(err)
	}
}
