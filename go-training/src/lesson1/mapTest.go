package main

import "fmt"

func main() {
	var map1 map[string] string = map[string] string {"key":"value"}
	fmt.Println(map1)
	fmt.Println(map1["key"])
	//fmt.Println(map1.key) wrong syntax

	map2 :=make(map[string] string,5)
	map2["1"] = "first"
	map2["2"] = "second"
	fmt.Println(map2)

	value,ok := map2["2"]
	if ok{
		fmt.Println(value +" was found")
	}

	sum:=0
	// just like while
	for {
		sum++
		fmt.Println(sum)
		if sum > 10 {
			break
		}
	}
}
