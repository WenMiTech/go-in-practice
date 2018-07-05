package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Phone string `json:"phone"`
}
type User struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Profile *Profile `json:"profile"`
}

func main() {
	user := &User{12, "test", &Profile{"13112271516"}}
	data, err := json.Marshal(user)
	if err == nil {
		fmt.Printf("%s", data)
	}

	user = &User{}
	err = json.Unmarshal([]byte(data), user)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(user.Profile)
}
