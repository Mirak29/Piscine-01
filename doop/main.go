package main

import (
	"os"
)

func isnum(s string) bool {
	for i, v := range s {
		if i == 0 && (v == '+' || v == '-') {
			for _, k := range s[1:] {
				if k < '0' || k > '9' {
					return false
				}
			}
		} else if i == 0 && (v != '+' || v != '-') {
			for _, k := range s {
				if k < '0' || k > '9' {
					return false
				}
			}
		}
	}
	return true
}
func inter(s string) int {
	a := 0
	if !isnum(s) {
		return 0
	}
	for _, v := range s {
		if v == '+' || v == '-' {
			continue
		}
		a *= 10
		a = a + int(v) - 48
	}
	if s[0] == '-' {
		a *= -1
	}
	return a
}
func stringer(s int) string {
	if s == 0 {
		return "0"
	}
	a := ""
	out := ""
	signe := ""
	if s < 0 {
		s *= -1
		signe = "-"
	}
	for s > 0 {
		a += string((s % 10) + 48)
		s /= 10
	}
	for i := len(a) - 1; i >= 0; i-- {
		out += string(a[i])
	}
	return signe + out
}
func main() {
	arg := os.Args
	if len(arg) != 4 {
		return
	}
	a, b := arg[1], arg[3]
	operator := arg[2]
	if !isnum(a) || !isnum(b) || ((inter(a) < -2147483647 && inter(a) > 2147483647) || (inter(b) < -2147483647 && inter(b) > 2147483647)) || (operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%") {
		return
	} else {
		resultat := 0
		switch operator {
		case "+":
			resultat = inter(a) + inter(b)
		case "-":
			resultat = inter(a) - inter(b)
		case "*":
			resultat = inter(a) * inter(b)
		case "/":
			if inter(b) == 0 {
				os.Stdout.WriteString("No division by 0\n")
				return
			}
			resultat = inter(a) / inter(b)
		case "%":
			if inter(b) == 0 {
				os.Stdout.WriteString("No modulo by 0\n")
				return
			}
			resultat = inter(a) % inter(b)
		}
		if resultat >= -2147483647 && resultat <= 2147483647 {
			os.Stdout.WriteString(stringer(resultat) + "\n")
		}
	}
}
