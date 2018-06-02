package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutines")

	go func(){
		defer wg.Done()
		for count:=0;count < 3;count++{
			for char:= 'a';char<'a'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()

	go func(){
		defer wg.Done()
		for count:=0;count < 3;count++{
			for char:= 'A';char<'A'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()

	 fmt.Println("waiting for finish")
	 wg.Wait()
	 fmt.Println("\n Terminating Program")
}
