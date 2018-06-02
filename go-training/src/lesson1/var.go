package main

import "fmt"

func main() {

	//常量
	const limit = 512
	const num int = 43
	//declare multi constance
	const (
		a = 0
		b = 2
		c = 4
	)

	// how to using iota
	const (
		aa = iota
		bb
		cc
	)
	// now iota's value is 2,then keyword `const` appear again,iota wiil reset to 0
	const dd =6
	const ff = iota
	fmt.Println(ff)//ff should be 0

	fmt.Println(aa,bb,cc)

	var isFlag bool = true
	isCreated := false

	fmt.Println(isFlag)
	fmt.Println(isCreated)

}
