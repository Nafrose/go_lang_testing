package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func revers(s string) string {
	//s = string(s[:])
	len := len(s)
	var r = make([]byte, len)

	for i := 0; i < len; i++ {
		r[i] = s[len-1-i]
	}
	result := string(r)

	return result
}

func main() {
	lines, err := readLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	strLines := fmt.Sprint(lines)
	rvrStrLines := revers(strLines)

	if err := writeLines(rvrStrLines, "output"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
