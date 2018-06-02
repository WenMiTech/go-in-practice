package person

type Person struct {
	Name string
	id   int
}

func New() *Person {
	return &Person{"mike", 35}
}
