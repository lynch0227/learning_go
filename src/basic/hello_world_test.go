package basic

import (
	"fmt"
	"os"
	"testing"
)

func variables() {
	aa, bb := 1, ""
	fmt.Printf("%d %q\n", aa, bb)
}

func TestHelloWorld(t *testing.T) {
	variables()
	if len(os.Args) > 1 {
		t.Log("Hello World!", os.Args[1])
	}
}
