package main

import "fmt"

func main() {

	fmt.Println(reverse("123456"))

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
