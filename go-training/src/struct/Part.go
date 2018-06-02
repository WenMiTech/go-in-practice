package main

import (
	"fmt"
	"strings"
)

type Part struct {
	Id   int
	Name string
}

func (part *Part) LowerCase() {

	part.Name = strings.ToLower(part.Name)

}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}
func main() {
	part := Part{3, "test"}
	fmt.Println(part)
	part.UpperCase()
	//&part.UpperCase() //complie error

	fmt.Println(part)

}
