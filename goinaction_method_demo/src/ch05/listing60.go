package main

import "fmt"

type notifer interface {
	notify()
} 

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email);
}

type admin struct {
	user
	level string
}

func (a *admin) notify()  {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name, a.email)
}

func sendNotification(n notifer) {
	n.notify()
}

func main() {

	ad := admin{
		user: user{
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	sendNotification(&ad)
	ad.user.notify()
	ad.notify()


}
