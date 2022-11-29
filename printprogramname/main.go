package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for i, v := range arguments {
		if i > 1 {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
