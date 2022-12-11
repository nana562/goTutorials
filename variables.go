package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	var (
	//		name string = "User"
	//		age int = 45
	//		location string = "Ghana"
	//	)
	//fmt.Println(name)
	//fmt.Println(location)
	//fmt.Println(age)
	//var i int
	i = 3
	//var j int = 23
	k := 45
	fmt.Println(i)

	var i int = 48
	fmt.Printf("%v, %T\n", i, i)

	var j int = 23
	var k float32
	fmt.Printf("%v, %T\n", k, k)
	k = float32(j)

	var s string
	s = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s, s)
}
