package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	tab := []rune(s)
	a := 0
	if len(tab) == 0 {
		return 0
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == 45 || tab[i] == 43 || (tab[i] >= 48 && tab[i] <= 57) {
			if i != 0 && (tab[i] == 45 || tab[i] == 43) {
				return 0
			}
			if tab[i] >= 48 && tab[i] <= 57 {
				a *= 10
				a = a + int(tab[i]) - '0'
			}

		} else {
			return 0
		}
	}
	if tab[0] == 45 {
		a *= -1
	}

	return a
}

func alphabet() []rune {
	var tab []rune
	for i := 'a'; i <= 'z'; i++ {
		tab = append(tab, i)
	}
	return tab
}

func ToUpper(s []rune) []rune {
	for i, v := range s {
		if v >= 97 && v <= 122 {
			s[i] = 65 + v - 97
		}
	}
	return s
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		return
	}
	alpha := alphabet()
	for i, v := range arguments {
		s := Atoi(v)
		if s == 0 && i == 0 {
			ToUpper(alpha)
		}
		if s > 0 && s < 27 {
			z01.PrintRune(alpha[s-1])
		}
		if ((s < 1 || s > 26) && v != "--upper") || (i != 0 && v == "--upper") {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
}
