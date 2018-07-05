package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	// "reflect"
)

type ResMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageView struct {
	id         int
	count      int
	createDate string
	updateDate string
}
type User struct {
	Name string `json:"name"`
}

func PicHandler(w http.ResponseWriter, r *http.Request) {

	user := &User{"test"}
	resMsg := &ResMsg{200, "success", user}
	data, err := json.Marshal(resMsg)
	if err != nil {
		fmt.Println(err)
	}
	header := w.Header()
	//header set must be called before w.writerheader
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(header)
	fmt.Fprintln(w, string(data))
}
func pvHandler(w http.ResponseWriter, r *http.Request) {
	dataSourceName := "root:root@tcp(localhost:3306)/found"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := db.Query("select *from page_views")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var id int
		var count int
		var createDate string
		var updateDate string
		err := rows.Scan(&id, &count, &createDate, &updateDate)
		if err != nil {
			log.Fatal("scan err", err)
			return
		}
		fmt.Printf("%d %d %s %s\n", id, count, createDate, updateDate)
	}
	defer db.Close()
	fmt.Fprint(w, "test")
}

func main() {
	http.HandleFunc("/goweb/pic", PicHandler)
	http.HandleFunc("/goweb/pv", pvHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
