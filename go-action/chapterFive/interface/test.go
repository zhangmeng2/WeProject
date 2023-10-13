package main

import "fmt"

type notifier2 interface {
	notify()
}

type user2 struct {
	name  string
	email string
}

func (u user2) notify() {
	fmt.Print("Sending user email to %s <%s>", u.name, u.email)
}

func main2222() {
	u := user2{"Bill", "bill@email.com"}
	sendNotification2(u)
}

func sendNotification2(n notifier) {
	n.notify()
}
