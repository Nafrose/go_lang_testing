package main

import (
	"fmt"
	"strings"
)

func cashier2(askingMoney int) string {
	notesPossibilities := [9]int{1000, 500, 100, 50, 20, 10, 5, 2, 1}
	remainingMoney := askingMoney
	printStacks := make([]string, 0, len(notesPossibilities))

	for _, note := range notesPossibilities {
		remainingMoney, printStacks = getRemainingAmount(remainingMoney, remainingMoney, note, printStacks)

		if remainingMoney <= 0 {
			break
		}
	}

	allPrintNotes := strings.Join(printStacks, ",\n")

	return fmt.Sprintf("{\n%s,\n\"requestAmount\":%d \n}", allPrintNotes, askingMoney)
}

func getRemainingAmount(askingMoney int, remainingMoney int, note int, printStacks []string) (int, []string) {
	givenNotes := remainingMoney / note
	// fmt.Print(givenNotes)

	if givenNotes > 0 {
		printNote := fmt.Sprintf("\"%d\":%d", note, givenNotes)
		printStacks = append(printStacks, printNote)
		remainingMoney = askingMoney - (givenNotes * note)
	} else {
		remainingMoney = askingMoney
	}

	return remainingMoney, printStacks
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
	printStackAsJson := cashier2(1503)
	fmt.Println(printStackAsJson)
	printStackAsJson2 := cashier2(28)
	fmt.Println(printStackAsJson2)
	// cashier(1.17)
}
