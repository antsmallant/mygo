package main

import (
	"fmt"
)


func main() {
	s := []int{1,2,3,4,5,6}
	ps(s)

	s = s[:4]
	ps(s)

	s = s[:6]
	ps(s)
}

func ps(s []int) {
	fmt.Printf("%v, len:%d, cap:%d\n", s, len(s), cap(s))
}