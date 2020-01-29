package basic

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string: " + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(i + 10)
	}
}
