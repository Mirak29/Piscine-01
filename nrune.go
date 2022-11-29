package piscine

func NRune(s string, n int) rune {
	tab := []rune(s)
	if len(tab) != 0 {
		for i, v := range tab {
			if i == n-1 {
				return v
			}
		}
	}
	return 0
}
