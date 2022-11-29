package piscine

func RecursivePower(nb int, power int) int {
	pow := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	pow *= nb
	if power > 1 {
		return nb * RecursivePower(nb, power-1)
	}
	return pow
}
