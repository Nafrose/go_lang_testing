package main

import (
	"io/ioutil"
)

func reverse(s []byte) []byte {
	len := len(s)
	var r = make([]byte, len)

	for i := 0; i < len; i++ {
		r[i] = s[len-1-i]
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
