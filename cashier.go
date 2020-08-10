package main

import "fmt"

func cashier2(given int) {
	notesPossbilities := [9]int{1000, 500, 100, 50, 20, 10, 5, 2, 1}
	

}

func cashier(t float64) {
	tInt := int(t * 100)

	q := tInt / 25
	d := (tInt - 25*q) / 10
	n := (tInt - 25*q - 10*d) / 5
	p := tInt - 25*q - 10*d - 5*n

	var m int = q + d + n + p

	fmt.Printf("The cashier(%d)\n", t)
	fmt.Printf("may give %d quarte(s)\n %d dime(s)\n %d nickel(s) \n %d penny(ies) which totals to %d coins", q, d, n, p, m)
}

func main() {
	cashier(1500)
	cashier(1.17)
}
