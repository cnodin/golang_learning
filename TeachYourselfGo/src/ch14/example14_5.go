package main

import (
	"errors"
	"fmt"
)

//Animal specifies an animal
type Animal struct {
	Name string	//Name holds the name of a thing
	//Age holds the name of a thing
	Age int
}

//ErrNotAnAnimal is returned if the name field of the Animal struct is Human
var ErrNotAnAnimal = errors.New("Name is not an animal")

func (a Animal) Hello() (string, error) {
	if a.Name == "Human" {
		return "", ErrNotAnAnimal
	}

	s := "Hello " + a.Name
	return s, nil
}

func main() {
	animal := new(Animal)
	animal.Age = 123
	animal.Name = "dog"

	s, _ := animal.Hello()
	fmt.Println(s)
}
