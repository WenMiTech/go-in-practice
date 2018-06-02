package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println(start)
	var a int
	for i := 0; i < 1000000000; i++ {
		a = a + 1
	}
	spendTime := time.Now().Sub(start)
	fmt.Println(spendTime)
	fmt.Println(a)
}
