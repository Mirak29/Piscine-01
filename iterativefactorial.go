package piscine

func IterativeFactorial(nb int) int {
	if nb >= 21 || nb <= -21 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	for i := nb - 1; i >= 1; i-- {
		nb *= i
	}
	return nb
}
