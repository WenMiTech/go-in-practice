package main2

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("client")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9091")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println(conn)
	fin, err := os.Open("test.txt")
	if err != nil {
		panic("open file err")
	}
	finfo, err := fin.Stat()
	if err != nil {
		fmt.Println("get file info err")
	}
	fsize := finfo.Size()

	buf := make([]byte, 1024)
	times := 0
	for {
		n, err := fin.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		n2, err := conn.Write(buf[:n])

		if err != nil {
			panic(err)
		}
		times += 1
		percent := float64(int64(times*1024) / fsize)
		fmt.Println(" %d send", percent)
		fmt.Println(n2)

	}
}
