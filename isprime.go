package piscine

func IsPrime(nb int) bool {
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		} else {
			continue
		}
	}
	if nb <= 0 || nb == 1 {
		return false
	}
	return true
}
