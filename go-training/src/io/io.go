
package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	for i := 0; i < 10; i++ {
		fout.WriteString("just test!\r\n")
		fout.Write([]byte("just test!\r\n"))
	}
	ReadFully()
}

func ReadFully() {
	userFile := "test.txt"
	in, err := os.Open(userFile)
	defer in.Close()
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
