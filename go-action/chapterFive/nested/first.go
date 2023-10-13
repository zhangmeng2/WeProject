package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  //嵌入类型
	level string
}

func main22() {
	ad := admin{
		user: user{
			name:  "john",
			email: "john@email.com",
		},
		level: "super",
	}

	ad.user.notify()

	ad.notify()
}
