package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		a := l.Head
		for a.Next != nil {
			a = a.Next
		}
		return a.Data
	}
}
