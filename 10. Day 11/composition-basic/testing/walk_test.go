package unittesting

import "testing"

func TestWalk(t *testing.T) {

	result := walk(3)

	if result != "Walking with 3 legs" {
		t.Errorf("Walk (3) failed with return of %v expected %v", "Walking with 3 legs", result)
	} else {
		t.Logf("Walk (3) succeeded with return of %v expected %v", "Walking with 3 legs", result)
	}
}
