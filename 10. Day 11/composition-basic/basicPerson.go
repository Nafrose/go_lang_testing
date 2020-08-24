package main

import "fmt"

type PersonBehavior interface {
	Walk()
	Talk()
}

type PersonProperties struct {
	hasHair bool
	legs    int8
	arms    int8
}

//Redundant structs Person and PersonProperties created just for observational purpose
type Person struct {
	PersonProperties
}

func (p Person) Walk() {
	fmt.Printf("Walking with %v legs.\n", p.legs)
}

func (p Person) Talk() {
	fmt.Printf("Talking now and gesturing with %d hands. \n", p.arms)
}

func main() {
	var alexProp PersonProperties
	alexProp.legs = 2
	alexProp.arms = 2
	var alex PersonBehavior = Person{alexProp}

	alex.Walk()
	alex.Talk()
}
