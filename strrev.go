package piscine

func StrRev(s string) string {
	var rev string
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		rev = rev + string(rune(s[i]))
	}
	return rev
}
