package main

import (
	"fmt"
	_ "strconv"
	piscine ".."
)

func main() {
	// fmt.Println(strconv.Atoi("1234"))
	s := ""
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := piscine.Atoi(s)
	n2 := piscine.Atoi(s2)
	n3 := piscine.Atoi(s3)
	n4 := piscine.Atoi(s4)
	n5 := piscine.Atoi(s5)
	n6 := piscine.Atoi(s6)
	n7 := piscine.Atoi(s7)
	n8 := piscine.Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
