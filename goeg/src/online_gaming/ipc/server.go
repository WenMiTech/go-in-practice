package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string
	Body string
}

type Server interface {
	Name() string
	Handle(method string,params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer{
	return &IpcServer{server}
}



func(ipcServer IpcServer) Connect()chan string{
	session:=make(chan string,0)


	go func(c chan string){
		for   {
			request:=<-c
			if request == "CLOSE"{
				break
			}
			var req Request
			err:=json.Unmarshal([]byte(request),&req)
			if err!=nil{
				fmt.Println("error format")
			}

			resp := ipcServer.Handle(req.Method,req.Params)
			b,err:=json.Marshal(resp)
			c<-string(b)

		}
		fmt.Println("session closed")
	}(session)
	return session
}
