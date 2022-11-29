package piscine

func BasicJoin(elems []string) string {
	elem := ""
	for _, v := range elems {
		elem += v
	}
	return elem
}
