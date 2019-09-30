package main

import (
	"os"
	"fmt"
)

func PrintString(s string) {
	str := []rune(s)
	for _, val := range str {
		z01.PrintRune(val)
	}
}

func main() {
	PrintString(os.Args[0])
}
