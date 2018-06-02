package main

import (
	"flag"
	"fmt"
)

// run   go run src/flagtest/flagtest01.go  -p /home/lyonru/ test test2
func main() {

	path := flag.String("p","","")//先定义,后解析
	flag.Parse()// flag must be parsed
	fmt.Println("path:",*path)
	var cmd string= flag.Arg(0)
	fmt.Println(flag.NArg())
	var cmd2 string = flag.Arg(1)

	fmt.Printf("first cmd : %s\n",cmd)
	fmt.Printf("second cmd : %s\n",cmd2)
}