package piscine

import "fmt"

func PrintComb() {
	for i:=0; i<=9; i++ {
		for j:=i+1; j<=9; j++ {
			for k:=j+1; k<=9; k++ {
				if i == 7 && j == 8 && k == 9 {
					fmt.Printf("%d%d%d\n", i, j, k)
				} else {
					fmt.Printf("%d%d%d, ", i, j, k)
				}
			}
		}
	}
}

