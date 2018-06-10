package main

import (
	"fmt"
	"goroutine/ptest"
	"runtime"
)

func main() {

	ptest.Ptest()
	name := "Eric"

	go func() {
		fmt.Println("hello", name)
	}()

	name = "Harry"
	runtime.Gosched()

}
