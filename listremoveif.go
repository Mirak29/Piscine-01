package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	actual := l.Head
	var ante *NodeL
	for actual != nil {
		if actual.Data == data_ref {
			if ante == nil {
				l.Head = l.Head.Next
			} else {
				ante.Next = actual.Next
			}
		} else {
			ante = actual
		}
	}
	actual = actual.Next
}
