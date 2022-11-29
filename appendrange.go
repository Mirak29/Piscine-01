package piscine

func AppendRange(min, max int) []int {
	var tab []int
	if min > max {
		return tab
	}
	for min < max {
		tab = append(tab, min)
		min += 1
	}
	return tab
}
