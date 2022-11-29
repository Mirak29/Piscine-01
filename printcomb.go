package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= 55; i++ {
		for j := '1'; j <= 56; j++ {
			for k := '2'; k <= 57; k++ {
				if i < j && j < k {

					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i != 55 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}
