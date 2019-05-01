package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int)int{}
	m[1] = func(opt int) int { return opt}
	m[2] = func(opt int) int { return opt * opt}
	m[3] = func(opt int) int { return opt * opt * opt}

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true

	n := 3

	if mySet[n] {
		t.Logf("%d is existed", n)
	}else {
		t.Logf("%d is not existed", n)
	}

	mySet[3] = true
	t.Log(len(mySet))
	t.Log(mySet)

	delete(mySet, 2)
	t.Log(mySet)
}
