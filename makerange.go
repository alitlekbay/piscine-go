package piscine

func SliceLen(newStr []int) int {
	n := 0
	for _, _ = range newStr {
		n++
	}
	return n
}

func myAppend(Str []int, num int) []int {
	i := 0
	n := SliceLen(Str) + 1
	newStr := make([]int, n)
	for _, val := range Str {
		newStr[i] = val
		i++
	}
	newStr[i] = num
	return newStr
}

func MakeRange(min, max int) []int {
	var newStr []int

	for ; min < max; min++ {
		newStr = myAppend(newStr, min)
	}
	return newStr
}
