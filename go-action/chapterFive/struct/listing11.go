package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("Sending User Email to %s<%s>", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

	//底层值被修改了
	lisa.notify()
	bill.notify()
}
