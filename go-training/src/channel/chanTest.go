package main

import (
	"fmt"
	"time"
)

func main() {

	c:=make(chan  bool)

	for i:=0;i<5;i++{
		go func(n int) {
			<- c
			fmt.Println("get message:",n)
		}(i)
	}
	fmt.Println("broacast...")
	close(c)//通道关闭时,所有协程退出阻塞
	time.Sleep(time.Second * 1)

}
