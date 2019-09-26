package piscine

func toInt(n byte) int {
	if n == 1 {
		return (1)
	} else if n == 2 {
		return (2)
	} else if n == 3 {
		return (3)
	} else if n == 4 {
		return (4)
	} else if n == 5 {
		return (5)
	} else if n == 6 {
		return (6)
	} else if n == 7 {
		return (7)
	} else if n == 8 {
		return (8)
	} else if n == 9 {
		return (9)
	}
	return 0
}

func BasicAtoi2(s string) int {
	
	newstr := []byte(s)
	count := 0
	for _, _ = range newstr {
		count++
	}

	result := 0
	for i := 0; i < count; i++ {
		if newstr[i] < '0' || newstr[i] > '9' {
			return 0
		}
		result = result * 10 + toInt((newstr[i]) - '0')
	}
	return result
}
