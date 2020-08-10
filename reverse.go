package main

import (
	"golang_test/console"
	"fmt"
)

// https://bit.ly/33KAQ0Y | https://bit.ly/3a9Biae
func main() {
	s, _ := Getline("Please enter your text to reverse:")
	fmt.Print("entered : ")
	fmt.Print(s)
	var r1, r2 = reverse(s)
	fmt.Printf("input: %s\n reverse1: %s \n reverse2: %s \n", s, r1, r2)
}

func reverse(s string) (string, string) {
	length := len(s)
	var r = make([]byte, length)
	var r2 = make([]byte, length)
	mid := int(length / 2)

	for i := 0; i < length; i++ {
		r[i] = s[length-1-i]
		fmt.Printf("%c-%c\n", s[i], s[length-1-i])
	}

	for i2 := 0; i2 < mid; i2++ {
		lastIndex := length - 1 - i2
		r2[i2] = s[lastIndex]
		r2[lastIndex] = s[i2]
	}

	r2[mid] = s[mid]

	reverse := string(r)
	reverse2 := string(r2)

	return reverse, reverse2
}
