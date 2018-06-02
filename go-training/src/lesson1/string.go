package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "abcdefg"
	fmt.Println(strings.LastIndex(str, "b"))

	var str1 string = "Tony 斯塔克"
	fmt.Println("字符串长度:",len(str1))
	//fmt.Println(range(str1))
	var substri  string = str1[0:5]
	fmt.Println(substri)
	char := str1[0]
	fmt.Println(char)  //char = 84
	fmt.Printf("char is %c",char)
	for i:=0;i<len(str1);i++{
		fmt.Println(str1[i])
	}

	//字符串的形式遍历
	for i,ch:=range str1{
		fmt.Println(i,ch)
		fmt.Println(string(ch))
	}
}
