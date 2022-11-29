package piscine

func ListReverse(l *List) {
	actual := l.Head
	var prev *NodeL
	prev = nil
	for actual != nil {
		next := actual.Next
		actual.Next = prev
		prev = actual
		actual = next
	}
	x := l.Head
	l.Head = l.Tail
	l.Tail = x
}
