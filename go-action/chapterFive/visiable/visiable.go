package visiable

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Rights int
}
