package piscine

func ConcatParams(args []string) string {
	sentence := ""
	for i, v := range args {
		sentence += v
		if i < len(args)-1 {
			sentence += "\n"
		}
	}
	return sentence
}
