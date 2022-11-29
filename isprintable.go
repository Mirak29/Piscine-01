package piscine

func IsPrintable(s string) bool {
	tab := []rune(s)
	for _, v := range tab {
		if v >= 32 {
			continue
		} else {
			return false
		}
	}
	return true
}
