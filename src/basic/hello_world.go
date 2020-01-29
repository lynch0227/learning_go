package basic

import (
	"fmt"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args[1])
	}
}
