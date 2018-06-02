package main

import "fmt"

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}

type Intf interface {
	M1()
	M2()
}

type S struct {
	T
}

type W struct {
	*T
}

func main() {

	//主要看方法的接受者？
	t1 := T{"t1"}
	fmt.Println(t1)
	t1.M1()
	fmt.Println(t1)
	t1.M2() //自动转为&t1
	fmt.Println(t1)

	t2 := &T{"t2"}
	fmt.Println(t2)
	t2.M1() //M1接收者为值类型,取t2的值拷贝给M1,t2可以看做默认参数  &t2 => t2
	fmt.Println(t2)
	t2.M2()
	fmt.Println(t2)

	//  t3为值类型，赋值给t4的是值类型，那么t4.M2()修改Name的值的时候是没有影响t3的，所以没有意义
	//var t3 T = T{"t3"}
	//var t4 Intf = t3// cannot use t3 (type T) as type Intf in assignment:
	//t4.M1()

	fmt.Println("t4,t3 test..............")
	var t3 T = T{"t3"}
	var t4 Intf = &t3
	t4.M1()
	fmt.Println("interface t4 =", t4)
	fmt.Println("interface t3 =", t3)
	fmt.Println("t5 test..............")
	t5 := T{"t5"}
	s := S{t5}
	fmt.Println(s.Name)
	s.M1()
	fmt.Println(s.Name)
	s.M2() //s{&t5.M2()}===>s{t5:"name2"}
	fmt.Println(s.Name)
	fmt.Println(t5.Name) //expected t5  ,应为S嵌入的是t5的拷贝，s.M2()方法的接受者t5自动转为&t5

	//var intf2 Intf=s//s还是值类型，M2无法改变s.Name

	fmt.Println("t6 test.....")
	t6 := T{"t6"}        // t6 ===> "t6"
	s1 := S{t6}          // 拷贝 t6 的值作为 s1的属性
	var intf2 Intf = &s1 // intf2  被赋值为s1 的引用

	intf2.M1()
	fmt.Println(s1.Name)                       //experted "t6"
	fmt.Println("var t6  experted: ", t6.Name) //

	intf2.M2()
	fmt.Println("M2 调用后")
	fmt.Println(s1.Name)
	fmt.Println("var t6 experted ", t6.Name)

	fmt.Println("t7 test....")
	w2 := W{&T{"t77"}}
	fmt.Println(w2)
	t7 := T{"t7"}
	w := W{&t7}
	fmt.Println("w M1 调用前")
	fmt.Println(t7.Name)
	fmt.Println(w.Name)
	w.M1() //M1 的接受者为值类型，
	fmt.Println("w M1 调用后")
	fmt.Println(t7.Name)
	fmt.Println(w.Name)
	w.M2()
	fmt.Println("w M2 调用后")
	fmt.Println(t7.Name)
	fmt.Println(w.Name)

}
