package main

import (
	"fmt"
	piscine ".."
)

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "Hello World!"

	n := piscine.BasicAtoi(s)
	n2 := piscine.BasicAtoi(s2)
	n3 := piscine.BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
}