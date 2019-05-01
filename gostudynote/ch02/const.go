package main

const x1 ,y int = 123, 0x22
const s = "hello, world"
const s1 = '我'
const s2 = '你'

const (
	Sunday = iota
	Monday
)

func main() {

	println(x1, y)
	println(s)
	println(s1)
	println(s2)

}
