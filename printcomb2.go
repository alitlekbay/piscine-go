package piscine

import "fmt"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for l := 0; l <= 9; l++ {
				for m := j + 1; m <= 9; m++ {
					if i == 9 && j == 8 && l == 9 && m == 9 {
						fmt.Printf("%d%d ", i, j)
						fmt.Printf("%d%d\n", l, m)
					} else {
						fmt.Printf("%d%d ", i, j)
						fmt.Printf("%d%d, ", l, m)
					}
				}
			}
		}
	}
}