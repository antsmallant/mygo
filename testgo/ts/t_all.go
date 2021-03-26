package main

import (
	"fmt"
	"reflect"
	"strings"
)

type VV struct {
	a string
}

func (v *VV) seta(a string) {
	v.a = a
}

func (v *VV) geta() string {
	return v.a
}

func test_VV() {
	var v VV = VV{a:fmt.Sprintf("%v", 200)}
	fmt.Println(v.geta())

	v.seta("100")
	fmt.Println(v.geta())

	v1 := &VV{a:"aaa"}
	fmt.Println(v1.geta())

	fmt.Println(reflect.TypeOf(v).String())	

	fmt.Println(reflect.TypeOf(v1).String())	
}

func test_slice() {
	var a []int
	fmt.Println(a)
	if a == nil {
		fmt.Println("nil")
	}
}


func printSlice(s []int) {
	fmt.Printf("%v, len:%d, cap:%d\n", s, len(s), cap(s))
}

func test_make_slice() {
	v := make([]int, 0, 5)
	printSlice(v)

	v = v[:3]
	printSlice(v)

	v = v[cap(v)-1:cap(v)]
	printSlice(v)

	v1 := make([]int, 10)
	printSlice(v1)

	v1 = append(v1, 100, 101)
	printSlice(v1)	
}

func test_double_slice() {
	b := [][]string {
		[]string{"-","-","o"},
		[]string{"-","-","o"},
		[]string{"-","-","o","x"},
	}
	fmt.Println(b)

	v := [5]int{}
	fmt.Println(v, len(v), cap(v))
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy, dy)
	for i := 0; i < dy; i++ {
		s[i] = make([]uint8, dx, dx)

		for j := 0; j < dx; j++ {
			s[i][j] = uint8(i*j)
		}
	}
	return s
}

func test_pic() {
	p := Pic(8, 8)
	fmt.Println(p)
}

func test_map() {
	v := map[string]string {
		"200":"300",
	}
	v["100"] = "aaa"
	fmt.Println(v)
	if _, ok:=v["100"]; ok {
		fmt.Println("exists 100")
	}
	delete(v, "100")
	fmt.Println(v)

	var v1 map[string]VV
	v1 = map[string]VV {
		"x":{"Y"},
	}
	v1["y"] = VV{"bb",}
	fmt.Println(v1)

	var v2 = make(map[string]string)
	v2["3"] = "4"
	fmt.Println(v2)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sf := strings.Fields(s)
	for _, f := range sf {
		cnt := 0
		if v, ok := m[f]; ok {
			cnt = v
		}
		m[f] = cnt+1
	}
	return m
}

func main() {
	test_map()
}