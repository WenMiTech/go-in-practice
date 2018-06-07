package main

import (
	"fmt"
	"net/http"
)

type Gogo struct {
}

func (gogo Gogo) Route(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	path := req.URL.Path
}
func (gogo Gogo) run() {
	fmt.Println("running gogo...")
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		gogo.Route(resp, req)
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("server down :", err)
	}
}

func main() {

	gogo := &Gogo{}
	gogo.run()
}
