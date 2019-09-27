package main

import (
	"fmt"
	piscine ".."
)

func main() {
    for i:=1;i<1000;i++ {
		// piscine.RecursiveFactorial(i)
    	fmt.Println(i, piscine.RecursiveFactorial(i))
	}
}