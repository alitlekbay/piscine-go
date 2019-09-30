package main

import (
	"os"
	"github.com/01-edu/z01"
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
