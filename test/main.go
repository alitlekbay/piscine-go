package main

import (
	"fmt"
	piscine ".."
)

func main() {
    arg1 := 4
	arg2 := 3
	fmt.Println(piscine.IterativePower(arg1, arg2))

	fmt.Println(piscine.IterativePower(-1, 2))
	fmt.Println(piscine.IterativePower(-1, 8))
	fmt.Println(piscine.IterativePower(-4, 4))
}