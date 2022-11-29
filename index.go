package piscine

func Index(s string, toFind string) int {
	if len(s) < len(toFind) || ((len(s) == len(toFind)) && s != toFind) {
		return -1
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
