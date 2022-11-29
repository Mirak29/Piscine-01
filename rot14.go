package piscine

func Rot14(s string) string {
	tab := []rune(s)
	for i, v := range tab {
		if v >= 65 && v < 91 {
			if v+14 > 90 {
				tab[i] = v + 14 - 26
			} else {
				tab[i] += 14
			}
		}
		if v >= 97 && v < 123 {
			if v+14 > 122 {
				tab[i] = v + 14 - 26
			} else {
				tab[i] += 14
			}
		}
	}
	rot := ""
	for _, v := range tab {
		rot += string(v)
	}
	return rot
}
