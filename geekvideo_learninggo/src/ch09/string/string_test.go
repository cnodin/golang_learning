package string_test

import "testing"

func TestString(t *testing.T)  {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中华民族"
	t.Log(s)
	t.Log(len(s))

	r := []rune(s)
	t.Log(len(r))

	t.Log(r)
	t.Logf("中 unicode %x", r[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"

	for _, v := range s {
		t.Logf("%[1]c %[1]x", v)
	}
}
