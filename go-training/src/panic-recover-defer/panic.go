package main

import (
	"fmt"
)

func main() {

	lessThan(2, 1)

	// fmt.Println(lessThan(2, 1))
}

func lessThan(a int, b int) bool {
	if a < b {
		return true
	}

	catchEx := func() {
		fmt.Println("panic  raise")
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}

	defer catchEx() // recover 一定是在defer语句中调用
	panic("b is bigger than a")
}
