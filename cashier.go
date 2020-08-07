package main

import "fmt"

func cashier(t float64) {
	tInt := int(t * 100)

	q := tInt / 25
	d := (tInt - 25*q) / 10
	n := (tInt - 25*q - 10*d) / 5
	p := tInt - 25*q - 10*d - 5*n

	var m int = q + d + n + p

	fmt.Printf("The cashier may give %d quarte(s), %d dime(s), %d nickel(s) and %d penny(ies) which totals to %d coins", q, d, n, p, m)
}

func main() {
	cashier(1.17)
}
