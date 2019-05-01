package main

var x = 10;

func main() {
	println(&x, x)

	x := "abc"
	println(&x, x)
}
