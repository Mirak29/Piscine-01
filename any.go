package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return f(v)
		}
	}
	return false
}
