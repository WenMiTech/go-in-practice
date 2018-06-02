package main

import "fmt"

type User2 struct {
	Name string
}

func(user *User2) SetName(name string){
	user.Name = name
}

func (user User2)talk ()  {
	fmt.Println(user.Name+" is talking")
}


type Manager struct {
	User2
}

func main() {
	m:=Manager{User2{"manage"}}
	m.SetName("manage2")
	m.talk()
	fmt.Println(m)

}

