package unittesting

import "testing"

func TestTalk(t *testing.T) {

	result := talk(3)

	if result != "Talking now and gesturing with 3 hands. \n" {
		t.Errorf("talk (3) failed with return of %v expected %v", "Talking now and gesturing with 3 hands. \n", result)
	} else {
		t.Logf("talk (3) succeeded with return of %v expected %v", "Talking now and gesturing with 3 hands. \n", result)
	}
}
