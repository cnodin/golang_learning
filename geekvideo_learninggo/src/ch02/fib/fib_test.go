package main

import "testing"

func TestFibList(t *testing.T)  {
	var (
		a int = 1
		b = 1
	)

	t.Log(a)
	t.Log(b)

	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	//tmp := a
	//a = b
	//b = tmp

	a, b = b, a

	t.Log(a, b)
}
