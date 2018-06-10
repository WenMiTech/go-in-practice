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
func (teacher Teacher) teach() {
	fmt.Println("teacher teaching...")
}

// so Principal owns A Teacher
// that it will has an Eat method
// that it is a Ihuman too
type Principal struct {
	Teacher
	Title string
}

func (pricipal Principal) Power() {
	fmt.Println("pricipal has highest power in the school")
}

func (pricipal Principal) teach() {
	pricipal.Teacher.teach()
	fmt.Println("pricipal teaching...")
}

//implements of Eat method
//so teacher is a Ihuman
func (teacher Teacher) Eat() {
	fmt.Println("eating food")
}

func main() {

	var humam Ihuman = Teacher{"tom"}

	humam.Eat()

	if humam, ok := humam.(Ihuman); ok {
		humam.Eat()
	}

	humam2 := &Principal{Teacher{}, "principal"}
	//actually call teacher.teach(),unless
	// unless principal has teach implements
	humam2.teach()
}
