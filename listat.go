package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := 0
	n := l
	for n != nil {
		if pos == i {
			return n
		} else {
			n = n.Next
			i++
		}
	}
	return nil
}
