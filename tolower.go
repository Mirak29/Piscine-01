package piscine

func ToLower(s string) string {
	tab := []rune(s)
	for i, v := range tab {
		if v >= 65 && v <= 90 {
			tab[i] = 97 + v - 65
		}
	}
	return string(tab)
}
