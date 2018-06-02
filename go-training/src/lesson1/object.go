package main

import (
	"fmt"
)

//better uppercase
type User struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {
	var bill User = User{"bill", "bill@163.com", 12, true}
	var tom *User = new(User)//返回一个应用

	fmt.Println(bill)
	fmt.Println(&tom)//address
	fmt.Println(tom)//instance
}
