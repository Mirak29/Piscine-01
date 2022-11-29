package piscine

func BasicAtoi(s string) int {
	tab := []rune(s)
	a := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] > 48 || tab[i] <= 57 {
			a *= 10
			a = a + int(tab[i]) - '0'

		}
	}

	return a
}
