package main

import (
	"strconv"
)

func main() {

/*	a,b,c := 100, 0144, 0x64

	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#x\n", a, b, c)

	fmt.Println(math.MinInt8, math.MaxInt8)*/

	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)

	println(a,b,c)
	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("0x" + strconv.FormatInt(a, 16))

	a1 := 1
	a1++
	println(a1)
}
