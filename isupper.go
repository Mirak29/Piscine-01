package piscine

func IsUpper(s string) bool {
	tab := []rune(s)
	for _, v := range tab {
		if v >= 65 && v <= 90 {
			continue
		} else {
			return false
		}
	}
	return true
}
