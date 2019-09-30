package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
	fmt.Println(piscine.Index("k]7Nz2Nxr70/Z", "Nxr7"))
}
