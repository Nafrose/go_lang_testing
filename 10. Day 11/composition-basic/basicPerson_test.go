package main

import (
	"testing"
)

func TestPersonStructForWalk(t *testing.T) {
	var alex Person
	alex.legs = 2

	resultOfWalk := alex.Walk()
	if resultOfWalk == "Walking with 2 legs. \n" {
		t.Logf("walk(3) succeeded with return of %v and expected return of : %v \n", "Walking with 2 legs. \n", resultOfWalk)
	} else {
		t.Errorf("walk(3) succeeded with return of %v and expected return of : %v \n", "Walking with 2 legs. \n", resultOfWalk)
	}
}

func TestPersonStructForTalk(t *testing.T) {
	var alex Person
	alex.arms = 2

	resultOfTalk := alex.Talk()
	if resultOfTalk == "Talking now and gesturing with 2 hands. \n" {
		t.Logf("talk(2) succeeded with return of %v and expected return of : %v \n", "Talking now and gesturing with 2 hands. \n", resultOfTalk)
	} else {
		t.Errorf("talk(2) succeeded with return of %v and expected return of : %v \n", "Talking now and gesturing with 2 hands. \n", resultOfTalk)
	}
}
