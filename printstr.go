package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	newstr := []rune(str)
	for _, value := range newstr {
		z01.PrintRune(value)
	}
}
