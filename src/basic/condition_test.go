package basic_try

import (
	"testing"
)

func TestCondition(t *testing.T) {
	a := 1
	if a == 1 {
		t.Log("Odd")
	} else if a == 2 {
		t.Log("Even")
	} else {
		t.Log("Other")
	}
}