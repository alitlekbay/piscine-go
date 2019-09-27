package main

import (
	"fmt"
	piscine ".."
)

func main() {
    arg1 := 4
	arg2 := 3
	fmt.Println(piscine.RecursivePower(arg1, arg2))

	fmt.Println(piscine.RecursivePower(-1, 2))
	fmt.Println(piscine.RecursivePower(-1, 8))
	fmt.Println(piscine.RecursivePower(-4, 4))
	fmt.Println(piscine.RecursivePower(-7, 7))

	// IterativePower(-7, 7) == 1 instead of -823543
    // IterativePower(-3, 4) == 1 instead of 81
    // IterativePower(-6, 3) == 1 instead of -216
    // IterativePower(-8, 3) == 1 instead of -512
    // IterativePower(-3, 9) == 1 instead of -19683
    // IterativePower(-4, 8) == 1 instead of 65536
    // IterativePower(-8, 7) == 1 instead of -2097152
    // IterativePower(-2, 3) == 1 instead of -8
}
