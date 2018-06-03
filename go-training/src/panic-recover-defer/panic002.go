package main

import (
	"fmt"
)

func main() {
	defer func() {

		v := recover()
		fmt.Println(v) //experted nil
		if v != nil {
			fmt.Printf("Recover a panic. [index=%d]\n", v)
		}
	}()
	fetchDemo()
	fmt.Println("main  fun executed")
}

func fetchDemo() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recover a panic. [index=%d]\n", v) //拦截之后，main函数重新获得控制权
		}
	}()
	ss := []string{"A", "B", "C"}
	fmt.Printf("Fetch  the elements in %v one by one...\n", ss)
	fetchElement(ss, 0)
	fmt.Println("fetch   done")
}

func fetchElement(ss []string, index int) (element string) {
	if index >= len(ss) {
		fmt.Printf("Occur a panic! [index=%d]\n", index)
		panic(index) //fetchElement 调用结束，即将调用defer，然后被fetchDemo的recover拦截
	}

	fmt.Printf("Fetching the element ... [index=%d]\n", index)
	element = ss[index]
	//defer 语句在外围函数将要执行结束之后被执行
	defer fmt.Printf("the element is \"%s\".[index=%d]\n", element, index)
	fetchElement(ss, index+1)
	return
}
