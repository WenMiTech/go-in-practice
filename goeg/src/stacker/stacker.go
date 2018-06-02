package main

import "fmt"
import "stacker/stack"

func main() {
	var heyStack stack.Stack
	heyStack.Push("you")
	var item, err = heyStack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(item)
}
