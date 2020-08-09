package main

import "fmt"

func main() {
	s := ""
	fmt.Println("Please enter your text to revserse:")
	fmt.Scanf("%s", &s)
	fmt.Printf("entered : %s", s)
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	len := len(s)

	var r = make([]byte, len)

	for i := 0; i < len; i++ {
		r[i] = s[len-1-i]
	}

	reverse := string(r)

	return reverse
}
