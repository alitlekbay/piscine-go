package piscine

func StrRev(s string) string {
	count := 0
	newStr := []byte(s)

	for _, _ = range newStr{
		count++
	}

	resstr := []byte(s)
	i := 0
	count--
	for num := count; num >= 0; num-- {
		resstr[i] = newStr[num]
		i++
	}

	return string(resstr)
}
