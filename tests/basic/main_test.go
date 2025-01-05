package basic

import "testing"

func AddOne(num int) int {
	return num + 1
}

func TestAddOne(t *testing.T) {

	var (
		in  = 1
		out = 3
	)

	if AddOne(in) != out {
		t.Error("Expected 2")
	}
}
