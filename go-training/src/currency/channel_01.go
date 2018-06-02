package main

import "fmt"
//bug
func main() {


	chs := make([]chan int, 10)
	fmt.Println(len(chs))
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])

	}

	for i, ch := range (chs) {
		fmt.Println(i)
		<-ch
	}
}

func Count(ch chan int) {
	ch <- 1 //把1写进通道,当ch还有数据时这个操作是阻塞的
	fmt.Println("Counting...")
}
