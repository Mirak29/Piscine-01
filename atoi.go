package piscine

func Atoi(s string) int {
	tab := []rune(s)
	a := 0
	if len(tab) == 0 {
		return 0
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == 45 || tab[i] == 43 || (tab[i] >= 48 && tab[i] <= 57) {
			if i != 0 && (tab[i] == 45 || tab[i] == 43) {
				return 0
			}
			if tab[i] >= 48 && tab[i] <= 57 {
				a *= 10
				a = a + int(tab[i]) - '0'
			}

		} else {
			return 0
		}
	}
	if tab[0] == 45 {
		a *= -1
	}

	return a
}
