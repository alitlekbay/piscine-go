package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	prime := 1
	for i := 2; i < nb; i++ {
		if ((nb % i) == 0) {
			prime = 0
		}
	}
	if prime == 1 {
		return true
	} else {
		return false
	}
}
