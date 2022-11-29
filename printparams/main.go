package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, v := range arguments[1:] {
		for _, j := range v {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
