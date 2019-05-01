package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify()  {
	//fmt.Printf("Address of u : %X", &u)
	fmt.Printf("Sending Email To %s<%s>\n",
		u.name, u.email)
}

func (u *user) changeEmail(email string)  {
	fmt.Printf("Address of u : %X\n", &u)
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	//fmt.Printf("Address of bill : %X", &bill)
	bill.notify()

	lisa := user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	fmt.Printf("Address of u : %X\n", &bill)
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
