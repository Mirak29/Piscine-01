package main

import (
	"os"

	"github.com/01-edu/z01"
)

func sorter(table []string) {
	tampon := ""
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1] {
				tampon = table[j]
				table[j] = table[j+1]
				table[j+1] = tampon

			}
		}
	}
}

func main() {
	arguments := os.Args[1:]
	sorter(arguments)
	for _, v := range arguments[0:] {
		for _, j := range v {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
