package basic

import (
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if len(os.Args) > 1 {
		t.Log("Hello World!", os.Args[1])
	}
}
