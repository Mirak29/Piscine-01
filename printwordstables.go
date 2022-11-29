package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		for _, j := range v {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
