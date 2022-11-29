package piscine

func AlphaCount(s string) int {
	cont := 0
	tab := []rune(s)
	for _, v := range tab {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			cont++
		}
	}
	return cont
}
