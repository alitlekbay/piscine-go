package piscine

func StrLen(str string) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}

func isNumeric(i byte) bool {
	if i >= '0' && i <= '9' {
		return true
	} else {
		return false
	}
}

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

func Atoi(s string) int {
	i := 0
	res := 0
	sign := 1

	if StrLen(s) == 0 {
		return 0
	}
	if s[0] == 0 {
		return 0
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}
	for ; i < StrLen(s); i++ {
		if (isNumeric(s[i]) == false) {
			return 0
		}
		res = res * 10 + toInt(s[i] - '0')
	}
	return sign * res
}