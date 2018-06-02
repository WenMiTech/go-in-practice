package count

import "fmt"

type Count int

func (count *Count) Increment()  { *count++ }
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

func main() {
	var count Count
	i := int(count) // 强行类型转换？
	fmt.Println(i)  //expected 0
	fmt.Println(count)
	count.Increment()
	fmt.Println(count)
	count.Decrement()
	fmt.Println(count)
	isZero := count.IsZero()
	fmt.Println(isZero)
}
