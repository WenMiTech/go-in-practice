package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {

	ad := admin{user: user{name: "lyonru", email: "lyonru@163,.com"}, level: "super"}
	ad.user.notify()
	ad.notify()//内部提升,接受者还是user
}
