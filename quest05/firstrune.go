package piscine

func StrLen(str string) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}

func FirstRune(s string) rune {
	ru := []rune(s)
	if StrLen(s) == 0 {
		return 0
	}
	return ru[0]
}
