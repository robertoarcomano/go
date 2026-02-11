package main

type user struct {
	name string
}

func (u *user) getName() string {
	return u.name
}

func (u *user) setName(name string) {
	u.name = name
}
