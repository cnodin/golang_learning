package main

import "fmt"

type notifier interface {
	notify()
}

type user1 struct {
	name string
	email string
}

func (u user1) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",
				u.name,
				u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user1{"Bill", "bill@email.com"}
	sendNotification(u)
}
