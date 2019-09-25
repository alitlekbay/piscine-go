package piscine

import "github.com/01-edu/z01"

func formatint(i int) rune {
	if i == 1 {
		return rune('1')
	} else if i == 2 {
		return rune('2')
	} else if i == 3 {
		return rune('3')
	} else if i == 4 {
		return rune('4')
	} else if i == 5 {
		return rune('5')
	} else if i == 6 {
		return rune('6')
	} else if i == 7 {
		return rune('7')
	} else if i == 8 {
		return rune('8')
	} else if i == 9 {
		return rune('9')
	} else if i == 0 {
		return rune('0')
	}
	return rune('0')
}

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for l := 0; l <= 9; l++ {
				for m := j + 1; m <= 9; m++ {
					if i == 9 && j == 8 && l == 9 && m == 9 {
						z01.PrintRune(formatint(i))
						z01.PrintRune(formatint(j))
						z01.PrintRune(rune(32))
						z01.PrintRune(formatint(l))
						z01.PrintRune(formatint(m))
						z01.PrintRune(rune(10))
					} else {
						z01.PrintRune(formatint(i))
						z01.PrintRune(formatint(j))
						z01.PrintRune(rune(32))
						z01.PrintRune(formatint(l))
						z01.PrintRune(formatint(m))
						z01.PrintRune(rune(44))
						z01.PrintRune(rune(32))
					}
				}
			}
		}
	}
}