package   main

import (
	"struct/department"
	
	"fmt"
)



type  PC struct {
	department.Department
}

func main(){

	pc:=PC{}
	fmt.Println(pc)
}