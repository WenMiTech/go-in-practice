package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"fmt"
)

var (
	url = "http://127.0.0.1:8090/download"
)

func main() {
	for i := 0; i < 1000000; i++ {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		path := "/Users/LyonRu/code/go/code/hellogo/src/test1/qq" + strconv.Itoa(i)+".txt"
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		io.Copy(f, res.Body)
		fmt.Println("file :",i, "downloaded")
	}

}