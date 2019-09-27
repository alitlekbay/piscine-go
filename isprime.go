package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	prime := 1
	for i := 2; i < nb; i++ {
		if ((nb % i) == 0) {
			return false
		}
	}
	if prime == 1 {
		return true
	} else {
		return false
	}
}
