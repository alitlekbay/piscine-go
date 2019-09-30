package piscine

// import "fmt"

func StrLen(str string) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}

func Index(s string, toFind string) int {
	match := -1
	lol := []rune(toFind)
	for k, val := range s {
		if k == StrLen(s) {
			break
		}
		if val == lol[0] {
			match = k
		} else {
			continue
		}
	}
	return match
}
