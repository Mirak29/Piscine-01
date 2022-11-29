package piscine

func Compact(ptr *[]string) int {
	i := 0
	for _, v := range *ptr {
		if v != "" {
			i++
		}
	}
	tab := make([]string, i)
	a := 0
	for _, v := range *ptr {
		if v != "" {
			tab[a] = v
			a++
		}
	}
	*ptr = tab
	return i
}
