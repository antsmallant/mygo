package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
	t.S = "100"
}

func main() {
	var i I
	i = &T{"hello"}
	i.M()
	fmt.Println(i)

	v := T{"World"}
	v.M()
	fmt.Println(v)
}