package main

import "fmt"

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Println("this is in admin")
}

func main() {

	ad := admin{
		user: user {
			name: "john",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	ad.user.notify()
	ad.notify()
}
