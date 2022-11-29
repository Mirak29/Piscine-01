package piscine

import (
	"github.com/01-edu/z01"
)

func printable(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func printrev(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		v := s[i]
		z01.PrintRune(rune(v))
	}
}

func validebase(base string) bool {
	compteneur := ""
	for i, v := range base {
		compteneur += string(v)
		for _, k := range base[i+1:] {
			if v == k {
				return false
			}
		}
	}
	return true
}

func corr(entier int, base string) string {
	convert := base[entier]
	return string(convert)
}

func PrintNbrBase(nbr int, base string) {
	signe := ""
	if nbr < 0 {
		nbr *= -1
		signe = "-"
	}
	signe = signe
	if !validebase(base) {
		printable("NV")
		return
	}
	hex := 0
	resultat := ""

	for nbr > 0 {
		hex = nbr % len(base)
		resultat += corr(hex, base)
		nbr /= len(base)
	}
	printrev(resultat + signe)
}
