package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// user 实现了 notifier 接口
func (u *user) notify() {
	fmt.Print("Sending user email to %s <%s>", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

// admin 实现了 notifier 接口
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>", a.name, a.email)
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

func sendNotification(n notifier) {
	n.notify()
}
