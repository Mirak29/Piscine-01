package piscine

func IsAlpha2(s rune) bool {
	if s > 32 {
	} else {
		return false
	}
	return true
}

func SplitWhiteSpaces(s string) []string {
	mot := ""
	var phrase []string
	for i, v := range s {
		if IsAlpha2(v) {
			mot = mot + string(v)
		} else if i != 0 && mot != "" {
			phrase = append(phrase, mot)
			mot = ""
		}
		if i == len(s)-1 {
			phrase = append(phrase, mot)
		}
	}
	return phrase
}
