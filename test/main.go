package main

import (
	"fmt"
	piscine ".."
)

func main() {
	for i := 1; i <= 1001; i++ {
		fmt.Println(piscine.IsPrime(i))
	}
}
