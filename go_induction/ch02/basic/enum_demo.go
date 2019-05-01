package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

func main() {
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
}
