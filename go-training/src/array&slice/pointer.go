package main

import "fmt"

/*
*
 */
func main() {
	var i int = 1
	fmt.Println(&i)
	var p *int = &i
	fmt.Println(*p) //experted *p = 1

	slice := []int{1, 2}  //slice
	arr := [2]int{12, 45} //arr
	modifySlice(slice)
	modifyArr(arr)
	fmt.Println(arr)
	fmt.Println(slice)
}

// srcArr is a slice
func modifySlice(srcSlice []int) {
	srcSlice[0] = 3
}

func modifyArr(arr [2]int) {
	arr[1] = 4
}
