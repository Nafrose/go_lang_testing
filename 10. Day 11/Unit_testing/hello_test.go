package Unit_testing

import "testing"

//test hello function with empty argument
func TestHelloWithEmptyArg(t *testing.T) {

	//test for empty argument
	emptyResult := hello("")

	if emptyResult != "Hello Dude!" {
		t.Errorf("Hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
	} else {
		t.Logf("Hello(\"\") success, expected %v, got %v", "Hello Dude!", emptyResult)
	}
}

//test hello function with valid argument
func TestHelloWithEmptyValidArg(t *testing.T) {

	//test for validy argument
	result := hello("Mike")

	if result != "Hello Mike!" {
		t.Errorf("Hello(\"\") failed, expected %v, got %v", "Hello Dude!", result)
	} else {
		t.Logf("Hello(\"\") success, expected %v, got %v", "Hello Dude!", result)
	}
}
