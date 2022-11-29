package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	a := l.Head
	for a != nil {
		if comp(a.Data, ref) {
			return &a.Data
		}
		a = a.Next
	}
	return &a.Data
}
