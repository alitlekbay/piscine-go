package main

import (
	"fmt"
	piscine ".."
)

func main() {
     for i:=1;i<101;i++ {
		// piscine.RecursiveFactorial(i)
    	fmt.Println(i, piscine.RecursiveFactorial(i))
	  }
}