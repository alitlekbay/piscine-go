package piscine

// import "fmt"

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb >= 1 {
		result := nb * RecursiveFactorial(nb - 1)
		if result < 0 {
			return 0
		}
		return result
	}
	return 1
}