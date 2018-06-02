package main

import (
	"fmt"
)

func main() {
	A()
	C()
	B()
}

func A() {
	fmt.Println("AAA")
}

func B() {

	fmt.Println("BBB")
}

func C() {
	defer C()
	panic("panic in c...")
}
