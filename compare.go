package piscine

func Compare(a, b string) int {
	tab1 := []rune(a)
	tab2 := []rune(b)
	less_long := 0
	if len(tab1) < len(tab2) {
		less_long = len(tab1)
	}
	if len(tab1) > len(tab2) {
		less_long = len(tab2)
	} else {
		less_long = len(tab1)
	}
	for i := 0; i <= less_long-1; i++ {
		if tab1[i] == tab2[i] {
			if i == less_long-1 && len(tab1) > len(tab2) {
				return 1
			} else if i == less_long-1 && len(tab1) < len(tab2) {
				return -1
			} else {
				continue
			}
		} else if tab1[i] < tab2[i] {
			return -1
		} else if tab1[i] > tab2[i] {
			return 1
		}
	}
	return 0
}
