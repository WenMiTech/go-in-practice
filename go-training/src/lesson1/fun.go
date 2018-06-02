package main

import "fmt"
import "lesson1/math"

func main() {
	fmt.Println(add(3,4))
	var c = math.Add(5,6)
	fmt.Println(c)
	paramFun(3,4,56)
	paramTest(3,65,6575)
}

//private
func add(a int,b int) int {
	return a + b
}

//不定参数
func paramFun(args ...int){
	for _,arg := range args{
		fmt.Println(arg)
	}
}

func paramTest(args ...int)  {
	test1(args...)
	test2(args[1:]...)
}

func test1(args ...int)  {
	fmt.Println("test1")
	for _,v := range args{
		fmt.Println(v)
	}
}

func test2(args ...int)  {
	fmt.Println("test2")
	for _,v := range args{
		fmt.Println(v)
	}
}