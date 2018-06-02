package main

import (
	"fmt"
)

type Item struct {
	id       string
	price    float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpectialItem struct {
	Item
	quantity  int
	catalogId int
}
type LuxuryItem struct {
	Item
	makeUp float64
}

func (item *LuxuryItem) Cost() float64 {
	// return item.Item.price * float64(item.Item.quantity) * item.makeUp
	// return item.price * float64(item.quantity) * item.makeUp
	return item.Item.Cost() * item.makeUp

}
func main() {
	item := Item{"1", 14.0, 3}
	cost := item.Cost()
	fmt.Println(cost)

	spectialItem := SpectialItem{Item{"2", 16.0, 4}, 5, 2}
	//有重名字段的访问情况
	fmt.Println("spectialItem.quantity  : ", spectialItem.quantity)
	fmt.Println("spectialItem.Item.quantity", spectialItem.Item.quantity)
	sCost := spectialItem.Cost()
	fmt.Println("sCost:", sCost)

	luxuryItem := LuxuryItem{Item{"3", 18.0, 5}, 3.0}
	lCost := luxuryItem.Cost()
	lItemCost := luxuryItem.Item.Cost()
	fmt.Println("lCost", lCost)
	fmt.Println("lItemCost", lItemCost)

}
