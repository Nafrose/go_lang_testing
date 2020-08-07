package main

import (
	"fmt"
	"math"
)

func main() {
	v := guess(5, []int{1, 2, 6, 7})
	fmt.Printf("The closest number is: %v", v)
}

func guess(n int, input []int) int {
	len := len(input)
	closest := 0
	winIndex := 0

	for i := 0; i < len; i++ {
		fDiff := math.Abs(float64(n - input[i]))
		diff := int(fDiff)
		if diff <= closest {
			closest = diff
			winIndex = i
		} else {
			if i == 0 {
				closest = diff
				winIndex = 0
			}
		}
	}

	return input[winIndex]
}
