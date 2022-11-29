package piscine

func ListSize(l *List) int {
	i := 0
	if l.Head == nil {
		return 0
	}
	for l.Head != nil {
		i++
		l.Head = l.Head.Next
	}
	return i
}
