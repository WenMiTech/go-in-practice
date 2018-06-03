package main

import (
	"fmt"
)

func main() {
	done := make(chan bool, 10)
	ch := make(chan int, 10)
	go produce(ch)

	go consumer(ch, done, "g1")
	go consumer(ch, done,"g2")
	//time.Sleep(time.Second *2)
	//<-done
	for i := 0; i < 10; i++ {
		<-done
	}
//        close(done)
	fmt.Print("done")
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
        close(p)

}

func consumer(c <-chan int, done chan bool,name string) {
	for v := range c {
		//v:=<-c
		fmt.Println(name, " receive:", v)
	}
	done <- true
}
