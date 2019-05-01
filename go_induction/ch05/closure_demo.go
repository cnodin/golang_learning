package main

func main() {
	str := "hello world"

	foo := func() {
		str = "hello dude"
	}

	foo()
}
