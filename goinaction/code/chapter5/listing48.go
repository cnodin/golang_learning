package main

import "fmt"

type notifier2 interface {
	notify()
}

type user2 struct {
	name string
	email string
}

func (u *user2) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := user{"lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

func sendNotification2(n notifier)  {
	n.notify()
}
