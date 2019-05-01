package const_test

import "testing"

const (
	Monday = iota + 1
	Thesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Eexecutable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Thesday, Wednesday)
}

func TestConstantTry1(t *testing.T)  {
	
}
