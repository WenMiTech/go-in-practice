package main2

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("server")

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9091")
	if err != nil {
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("listening...")
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("accepting...")
		go HandleConn(conn)
	}

}

func HandleConn(conn *net.TCPConn) {
	fmt.Println("handling the conn")
	fileName := time.Now().String() + ".txt"
	fout, err := os.Create(fileName)
	if err != nil {
		fmt.Println("err in create file")
	}
	defer fout.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			panic("read data error")
		}
		if n == 0 {
			break
		}
		if n2, err := fout.Write(buf[:n]); err != nil {
			panic("error in writing data")
		} else if n != n2 {
			panic("error in writing data")
		}
	}

}
