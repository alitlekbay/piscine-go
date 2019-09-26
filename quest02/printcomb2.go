package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int32) {
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

func PrintComb2() {
	var i int32
	var j int32
	 for i = 0; i <= 99; i++ {
	 	for j = 0; j <= 99; j++ {
	 		if i < j {
	 			if i < 10 {
	 				z01.PrintRune('0')
	 			}
	 			PrintNbr(i)
	 			z01.PrintRune(' ')
	 			if j < 10 {
	 				z01.PrintRune('0')
	 			}
	 			PrintNbr(j)
	 			if i == 98 && j == 99 {
	 				z01.PrintRune('\n')
	 			} else {
	 				z01.PrintRune(',')
	 				z01.PrintRune(' ')
	 			}
	 		}
	 	}
	}

}
