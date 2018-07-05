package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	s = "1"
	sep = "2"
	test := "3"
	s += sep + test
	fmt.Println(s)
}
