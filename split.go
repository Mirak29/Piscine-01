package piscine

func Split(s, sep string) []string {
	mot := []string(nil)
	cumulator := ""
	if len(sep) >= 1 {
		for i := 0; i < len(s); i++ {
			if (string(s[i]) != string(sep[0]) || s[i:i+len(sep)] != sep[0:]) && i != len(s)-1 {
				cumulator += string(s[i])
			} else {
				if i == len(s)-1 {
					cumulator += string(s[i])
				}
				if i != 0 && cumulator != "" {
					mot = append(mot, cumulator)
				}
				cumulator = ""
				i += len(sep) - 1
			}
		}
	}
	return mot
}
