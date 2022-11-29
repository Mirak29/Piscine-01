package piscine

func ToUpper(s string) string {
	tab := []rune(s)
	for i, v := range tab {
		if v >= 97 && v <= 122 {
			tab[i] = 65 + v - 97
		}
	}
	return string(tab)
}
