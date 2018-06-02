package main

import "fmt"

func main(){
	// the rest element of arr will be init to 0
	var arr [10]int = [10]int{1,2,3,4}
	for i,v:= range arr{
		fmt.Println(v,i)
	}

	//slice
	var mySlice []int = arr[:5]
	for i,v:= range mySlice{
		fmt.Println(v,i)
	}

	//other way to create a slice
	var slice2 []int = make([]int,5)
	fmt.Println(`slice2`)
	for i,v:= range slice2{
		fmt.Println(v,i)
	}
	//create a slice that has 5 element.and capacity is 10
	slice3 := make([]int,5,10)
	for i:=0;i<4;i++{
		slice3=append(slice3,i)
	}
	fmt.Println(slice3)

	//fun:copy
	slice4:=[]int{1,2,34,6,7}
	slice5:=[]int{4,5,6}
	copy(slice4,slice5)
	fmt.Println(slice4)//4,5,6,6,7
}
