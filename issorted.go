package piscine

func IsSorted(f func(int, int) int, a []int) bool {
	croissant, descroissant := true, true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			croissant = false
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			descroissant = false
		}
	}

	return croissant || descroissant
}
