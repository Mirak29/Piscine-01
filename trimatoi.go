package piscine

func TrimAtoi(s string) int {
	tab := []rune(s)
	a := 0
	minus := -1
	num := -1
	comp := 0
	compnum := 0
	if len(tab) == 0 {
		return 0
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] >= 48 && tab[i] <= 57 {

			a *= 10
			a = a + int(tab[i]) - '0'

			if compnum == 0 {
				num = i
				compnum++
			}

		}
		if tab[i] == 45 && comp == 0 {
			minus = i
			comp++
		}
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == 45 && a != 0 {
			if minus <= num {
				a *= -1
				return a
			}
		}
	}
	return a
}
