package main

import (
	"fmt"
	"os"
)

func main() {

	ReadFully()

}

func ReadFully() {
	userFile := "test1.txt"
	in, err := os.Open(userFile)

	// defer 的调用方式按先进后出的方式
	defer in.Close()
	defer func() {
		fmt.Println("defer testing...")
	}()
	defer func() {
		fmt.Println("defer testing 2...")
	}()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := in.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
