package basic

import (
	"errors"
	"fmt"
	"testing"
)

var ErrorBetweenHandle = errors.New("n should be in [2,100]")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, ErrorBetweenHandle
	}
	FibList := []int{1, 1}

	for i := 2; i < n; i++ {
		FibList = append(FibList, FibList[i-2]+FibList[i-1])
	}
	return FibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestPanicVsExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong"))
	// os.Exit(-1)
	// fmt.Println("End")
}
