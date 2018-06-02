package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
)

func helloHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(w.Header());
	io.WriteString(w,"hello,go http server")
}
func main()  {

	http.HandleFunc("/hello",helloHandler)
	err:=http.ListenAndServe(":8080",nil)
    if err!=nil{
	    log.Fatal("Listening for request...",err.Error())
    }
}
