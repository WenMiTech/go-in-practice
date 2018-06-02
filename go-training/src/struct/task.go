package main

import (
	"fmt"
	"struct/count"
	"struct/department"
	"struct/person"
)

type Task struct {
	slice []string
	count.Count
	person *person.Person
	department.Department
	id int
}

func main() {
	task := Task{
		[]string{"hah"},
		3,
		person.New(),
		department.Department{}, 34}
	fmt.Println(task.person.Name)
}
