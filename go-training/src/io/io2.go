package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fi, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("opening file")
	defer fi.Close()

	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buf := make([]byte, 1024)
	for {

		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if n2, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		} else if n2 != n {
			panic("error in writing")
		}

	}
}
