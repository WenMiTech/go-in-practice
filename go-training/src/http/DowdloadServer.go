package main

import (
	"net/http"
	"log"
	"os"
	"io"
	"path"
)

func main() {

	http.HandleFunc("/download", downloadHandler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("fail to start server :", err.Error())
	}
}

//实现文件下载
func downloadHandler(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("/Users/LyonRu/code/go/code/hellogo/src/http/readme.txt")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer file.Close()
	fileName := path.Base("/Users/LyonRu/code/go/code/hellogo/src/http/readme.txt")
	header := writer.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("content-disposition", "attachment; filename=\"" + fileName + "\"")
	io.Copy(writer, file)
}