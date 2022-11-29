package piscine

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
