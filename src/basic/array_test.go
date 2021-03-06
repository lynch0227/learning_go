package basic

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1[1], arr2[2])
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for _, e := range arr {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSec := arr[0]
	t.Log(arrSec)
}
