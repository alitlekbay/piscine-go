package piscine

func StrLen(str []rune) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}

func LastRune(s string) rune {
	sNew := []rune(s)
	return sNew[StrLen(sNew) - 1]
}
