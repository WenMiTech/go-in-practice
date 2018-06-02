package main

import (
	"fmt"
)

func main() {

	c := make(chan bool)
	go func() {
		fmt.Println("go go go")
		c <- true
	}()

	// 阻塞,直到c有值
	<-c
	testFunc()

}


func testFunc(){
	fmt.Println("test")
}