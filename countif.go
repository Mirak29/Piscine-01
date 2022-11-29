package piscine

func CountIf(f func(string) bool, tab []string) int {
	compt := 0
	for _, v := range tab {
		if f(v) == true {
			compt++
		}
	}
	return compt
}
