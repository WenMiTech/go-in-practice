package main

import (
	"fmt"
)

type Ihuman interface {
	Eat()
}

type Teacher struct {
	Name string
}

func (teacher Teacher) Print() {
	fmt.Println(teacher.Name)
}

func (teacher Teacher) Eat() {
	fmt.Println("eating food")
}
func main() {

	var humam Ihuman = Teacher{"tom"}
	humam.Eat()
	if humam, ok := humam.(Ihuman); ok {
		humam.Eat()
	}
}
