package piscine

func IsAlphaRune(s rune) bool {
	if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) || (s >= 48 && s <= 57) {
		return true
	}
	return false
}

func Capitalize(s string) string {
	blank := true
	tab := []rune(s)

	for i := 0; i < len(s); i++ {
		if blank && IsAlphaRune(tab[i]) {
			if tab[i] >= 'a' && tab[i] <= 'z' {
				tab[i] = tab[i] - 32
			}
			blank = false
		} else if tab[i] >= 'A' && tab[i] <= 'Z' {
			tab[i] = tab[i] + 32
		} else if !IsAlphaRune(tab[i]) {
			blank = true
		}
	}
	return string(tab)
}
