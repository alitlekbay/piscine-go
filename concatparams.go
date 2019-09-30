package piscine

func LenStr(args []string) int {
	count := 0
	for _, _ = range args {
		count++
	}
	return count
}

func Len(str string) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}

func ConcatParams(args []string) string {
	count := LenStr(args)
	num := 0
	for _, val := range args {
		num += Len(val)
	}
	counter := 0
	newStr := make([]byte, num + count - 1)
	for num, key := range args {
		for i := 0; i < Len(key); i++ {
			newStr[counter] = key[i]
			counter++
		}
		if num != LenStr(args) - 1 {
			newStr[counter] = 10
			counter++
		}
	}
	return string(newStr)
}
