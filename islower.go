package piscine

func IsLower(s string) bool {
	tab := []rune(s)
	for _, v := range tab {
		if v >= 97 && v <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
