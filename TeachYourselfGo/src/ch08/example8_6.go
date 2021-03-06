package main

import (
	"github.com/pkg/errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
	Talk()
}

type T850 struct {
	Name string
}

type R2D2 struct {
	Broken bool
}

func (a *T850) PowerOn() error {
	return nil
}

func (a *T850) Talk() {
	fmt.Println("T850 Talk")
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func (r *R2D2) Talk() {
	fmt.Println("R2D2 Talk")
}

func Boot(r Robot) error {
	r.Talk()
	return r.PowerOn()
}

func main() {
	t := T850{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

}
