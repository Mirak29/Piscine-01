package piscine

func ActiveBits(n int) int {
	compt := 1
	for n/2 != 0 {
		if n%2 == 1 {
			compt++
		}
		n /= 2
	}
	return compt
}
