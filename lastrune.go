package piscine

func LastRune(s string) rune {
	tab := []rune(s)
	return tab[len(tab)-1]
}
