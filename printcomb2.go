package piscine

import "github.com/01-edu/z01"

func formatint(i int) rune {
	if i == 1 {
		return ('1')
	} else if i == 2 {
		return ('2')
	} else if i == 3 {
		return ('3')
	} else if i == 4 {
		return ('4')
	} else if i == 5 {
		return ('5')
	} else if i == 6 {
		return ('6')
	} else if i == 7 {
		return ('7')
	} else if i == 8 {
		return ('8')
	} else if i == 9 {
		return ('9')
	} else if i == 0 {
		return ('0')
	}
	return ('0')
}

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for l := 0; l <= 9; l++ {
				for m := j + 1; m <= 9; m++ {
					if i == 9 && j == 8 && l == 9 && m == 9 {
						z01.PrintRune(formatint(i))
						z01.PrintRune(formatint(j))
						z01.PrintRune(32)
						z01.PrintRune(formatint(l))
						z01.PrintRune(formatint(m))
						z01.PrintRune(10)
					} else {
						z01.PrintRune(formatint(i))
						z01.PrintRune(formatint(j))
						z01.PrintRune(32)
						z01.PrintRune(formatint(l))
						z01.PrintRune(formatint(m))
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}