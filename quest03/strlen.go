package piscine

func StrLen(str string) int {
	count := 0
	newstr := []rune(str)
	for _, _ = range newstr {
		count++
	}
	return count
}