package main

import (
	"WeProject/go-action/chapterFive/visiable"
	"fmt"
)

func main() {
	a := visiable.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "Bill@Email.com"

	fmt.Printf("User :%v\n", a)
}
