package main

import "fmt"

func main() {
	a := 10
	aP := &a
	aPP := &aP

	fmt.Printf("a: %d\n", a)
	fmt.Printf("aP: %x\n", aP)
	fmt.Printf("*aP: %d\n", *aP)

	fmt.Printf("aPP: %x\n", aPP)
	fmt.Printf("*aPP: %x\n", *aPP)
}
