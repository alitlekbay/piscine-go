package piscine

func Index(s string, toFind string) int {
	lol := []rune(toFind)
	for k, val := range s {
		if val == lol[0] {
			return k
		}
	}
	return -1
}
