package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	//类似其他语言的while (n < 5)
	for n < 5 {
		t.Log(n)
		n++
	}
}
