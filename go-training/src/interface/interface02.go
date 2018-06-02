package main

import (
	"fmt"
)

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair, pair.first, pair.second)
}

type Point [2]int

func (point Point) Exchange() {

	point[0], point[1] = point[1], point[0]
}

func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func main() {
	jekyll := StringPair{"Herry", "Jekyll"}

	hyde := StringPair{"Edward", "Hyde"}
	point := Point{5, -3}

	fmt.Println("before", jekyll, hyde, point)

	jekyll.Exchange()
	hyde.Exchange()
	point.Exchange()
	// exchangeThese(jekyll)  ExChange need a pointer recevicer,but jekyll not

	exchangeThese(point)
}
