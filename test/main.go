package main

import (
	"fmt"
	piscine ".."
)

func main() {
     for i:=1;i<101;i++ {
          fmt.Println(i, piscine.IterativeFactorial(i))
	  }
}