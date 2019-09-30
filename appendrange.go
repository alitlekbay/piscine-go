package piscine

func AppendRange(min, max int) []int {
	var newStr []int

	for ; min < max; min++ {
		newStr = append(newStr, min)
	}
	return newStr
}
