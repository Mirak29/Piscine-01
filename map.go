package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, len(a))
	for i, v := range a {
		tab[i] = f(v)
	}
	return tab
}
