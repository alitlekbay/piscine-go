package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else if power != 0 {
		return nb * RecursivePower(nb, power -1)
	} else {
		return 1
	}
}
