package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = n*-1
	}
	if n >= 10 {
		PrintNbr(n/10)
	}
	z01.PrintRune(n % 10 + '0')
	return
}
