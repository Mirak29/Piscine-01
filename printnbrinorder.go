package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []int
	if n == 0 {
		tab = append(tab, 0)
	}
	for n > 0 {
		tab = append(tab, n%10)
		n = n / 10
	}
	SortIntegerTable(tab)
	for i := range tab {
		z01.PrintRune('0' + rune(tab[i]))
	}
}
