package main

import (
	"fmt"
)

type T struct {
	Name string
}

func (t *T) SetName(name string) {
	t.Name = name
}

func (t *T) SetName() {
	fmt.Println("can't overide")
}

func main() {
	t := T{"type1"}
	t.SetName("type2")
	fmt.Println(t)
}
