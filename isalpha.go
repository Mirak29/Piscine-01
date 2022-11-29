package piscine

func IsAlpha(s string) bool {
	tab := []rune(s)
	for _, v := range tab {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57) {
			continue
		} else {
			return false
		}
	}
	return true
}
