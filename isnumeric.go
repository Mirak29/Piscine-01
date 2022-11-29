package piscine

func IsNumeric(s string) bool {
	tab := []rune(s)
	for _, v := range tab {
		if v >= 48 && v <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
