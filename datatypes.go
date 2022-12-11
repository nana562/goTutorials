package main

import (
	"fmt"
)

func main() {
	//var n bool = false
	//var n bool = true
	//	fmt.Printf("%v, %T\n", n, n)
	a := 10
	//b := 7
	//var n complex64 = 1 + 3i
	//fmt.Println(a + b)
	//fmt.Println(a - b)
	//fmt.Println(a / b)
	//fmt.Println(a * b)
	//fmt.Println(a % b)
	//fmt.Printf("%v, %T\n", n, n)

	s := "this is a string"
	s1 := " this is also a string"
	b := []byte(s1)
	fmt.Printf("%v, %T\n", s+s1, s)
	fmt.Printf("%v, %T\n", b, b)

}
