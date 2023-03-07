package Testing_Go

import "testing"

func TestAddition(t *testing.T) {
	// testing go built in capability
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got %v, wanted %v", got, expected)
	}
}

//a test we expect to fail
func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 7
	if got != expected {
		t.Errorf("Did not get expected result. Got %v, wanted %v", got, expected)
	}
}
