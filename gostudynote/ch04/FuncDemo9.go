package main

import (
	"github.com/pkg/errors"
	"log"
)

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}

	return x / y, nil
}

func main() {
	z, err := div(5, 0)
	if err != nil {
		log.Fatalln(err)
	}

	println(z)

}
