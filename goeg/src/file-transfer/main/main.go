package main

import (
	"fmt"
	"lyonru/util"
	"net"
	"os"
)

func CheckError(err error, info string) (resp bool) {

	if err != nil {
		fmt.Println(info + "" + err.Error())
		return false
	}
	return true

}

func Handler(conn net.Conn, message chan string) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if CheckError(err, "connection") == false {
			conn.Close()
			break
		}
		if length > 0 {
			buf[length] = 0
		}

		reciveStr := string(buf[0:length])
		message <- reciveStr
	}
}

func echoHandler(conns *map[string]net.Conn, message chan string) {
	for {
		msg := <-message
		fmt.Println(msg)

		for key, value := range *conns {
			fmt.Println("connection is connected from .....", key)
			_, err := value.Write([]byte(msg))
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

func StartServer(port string) {
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err, "ListenTCP")

	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)

	go echoHandler(&conns, messages)

	for {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		CheckError(err, "Accept")
		fmt.Println("Accepting...")
		conns[conn.RemoteAddr().String()] = conn
		go Handler(conn, messages)
	}
}

func chatSend(conn net.Conn) {
	var input string
	username := conn.LocalAddr().String()
	for {
		fmt.Scanln(&input)
		if input == "/quit" {
			fmt.Println("bye bye ...")
			conn.Close()
			os.Exit(0)
		}

		lens, err := conn.Write([]byte(username + "Say:::" + input))
		fmt.Println(lens)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}
	}
}

func StartClient(tcpaddr string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	fmt.Println(tcpAddr)
	CheckError(err, "ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err, "DialTCP")
	go chatSend(conn)
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if CheckError(err, "Connection") == false {
			conn.Close()
			fmt.Println("Server is dead... ByeBye")
			os.Exit(0)

		}
		fmt.Println(string(buf[0:length]))
	}

}

func main() {
    
	util.PrintTime()
	util.Test()
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	
	if len(os.Args) != 3 {
		fmt.Println("exit....")
		os.Exit(0)
	}

	if os.Args[1] == "server" && len(os.Args) == 3 {
		StartServer(os.Args[2])
	}
	if os.Args[1] == "client" && len(os.Args) == 3 {
		StartClient(os.Args[2])
	}

}
