package main

import (
	"fmt"
)

func main() {

	//rune 字节？？ 汉字用3个字节表示？   3 * 4 = 12
	str1 := "暴殄天物" // len(str1) experted 12

	fmt.Println(len(str1))
	fmt.Println(str1[2:3])
	chars := []rune(str1) //码点切片  rune === int32
	strSlice := string(chars[:2])
	fmt.Println("中文字符串截取: ", strSlice)
	for _, char := range chars {
		fmt.Println(char)
		//fmt.Printf("[0x%X]  '%c'", char, char)
		fmt.Println(string(char)) //转换成字符
	}
}
