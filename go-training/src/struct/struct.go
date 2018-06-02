package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	B
	Age int
}

type D struct {
	B
	Name string
	Sex  int
}

func main() {
	var a = A{"aa"}
	fmt.Println(a)
	NameChange(a);
	fmt.Println(a)

	NameChange2(&a);
	fmt.Println(a)

	var aa = &A{"aaaa"}
	fmt.Println(aa)
	NameChange2(aa);
	fmt.Println(aa)
	//var b = B{"bb"}
	//fmt.Println(b)
	//
	//var c = C{Age: 12, B: B{Name: "cc"}}
	//fmt.Println(c)
	//
	//cc := C{Age: 14, B: B{Name: "cccc"}}
	//fmt.Println(cc)
	//fmt.Println(cc.Name)
	//
	//dd := D{B: B{Name: "testcccc"}, Name: "dddd", Sex: 1}
	//fmt.Println(dd)
	//fmt.Println(dd.B.Name)
	//dd.B.Name = "bbbbbbbb"
	//fmt.Println(dd.B.Name)
	//
	//newDDD := new(D)
	//newDDD.Name = "testDDDD"
	//fmt.Println(newDDD.Name)

}

// struct 是以值的形式传递，这里传递的是值的拷贝
func NameChange(a A) {
	a.Name = "aaaChange"
	fmt.Println("NameChange:", a)
}

// 要修改a的值，必须得用指针的形式传递
func NameChange2(a *A) {
	a.Name = "aaaChange2"
	fmt.Println("NameChange2:", a)
}