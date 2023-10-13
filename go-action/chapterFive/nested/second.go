package main

import "fmt"

type notifier3 interface {
	notify()
}

type user2 struct {
	name  string
	email string
}

func (u *user2) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin2 struct {
	user2 //嵌入类型
	level string
}

func (a *admin2) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin2{
		user2: user2{
			name:  "john",
			email: "john@email.com",
		},
		level: "super",
	}

	sendNotification(&ad)
	ad.user2.notify()
	ad.notify()
}

func sendNotification(n notifier3) {
	n.notify()
}
