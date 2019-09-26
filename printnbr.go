package piscine

import "github.com/01-edu/z01"

func PrintIt(n int) rune {
	if n == 1 {
		return ('1')
	} else if n == 2 {
		return('2')
	} else if n == 3 {
		return('3')
	} else if n == 4 {
		return('4')
	} else if n == 5 {
		return('5')
	} else if n == 6 {
		return('6')
	} else if n == 7 {
		return('7')
	} else if n == 8 {
		return('8')
	} else if n == 9 {
		return('9')
	}
	return '0'
}

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
	z01.PrintRune(PrintIt(n % 10))
	return
}
