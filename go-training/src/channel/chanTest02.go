package main

import (
	"fmt"
	"time"
)

func main() {

	ch:=make(chan int,10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(time.Second *2)


}

func produce(p chan<- int)  {
	for i:=0;i<10;i++{
		p <- i
		fmt.Println("send:",i)
	}

}


func consumer(c <-chan int){
	for i:=0;i<10;i++{
		v:=<-c
		fmt.Println("receive:",v)
	}
}


