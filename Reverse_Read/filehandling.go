package main

import (
	"io/ioutil"
)

func reverse(s []byte) []byte {
	var lenOf = len(s)
	var r = make([]byte, lenOf)

	for i := 0; i < lenOf; i++ {
		r[i] = s[lenOf-1-i]
	}

	reverse := r

	return reverse
}

func main() {
	// read the whole file at once
	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	text := reverse(b)

	//write the whole body at once
	err = ioutil.WriteFile("output", text, 0644)
	if err != nil {
		panic(err)
	}
}
