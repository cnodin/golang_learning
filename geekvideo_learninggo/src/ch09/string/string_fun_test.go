package string

import (
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")

	for _, p := range parts {
		t.Log(p)
	}

	t.Log(strings.Join(parts, "-"))

	t.Log(strings.ToLower(s))
}
