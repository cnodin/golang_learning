package main

import "fmt"

func main() {

	var chesses = make([]string, 3)
	chesses[0] = "Mariolles"
	chesses[1] = "Epoisses de Bourgogne"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, chesses)
	fmt.Println(smellyCheeses)
}
