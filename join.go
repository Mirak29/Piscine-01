package piscine

func Join(strs []string, sep string) string {
	elem := ""
	for i, v := range strs {
		if i != len(strs)-1 {
			elem += v
			elem += sep
		} else {
			elem += v
		}
	}
	return elem
}
