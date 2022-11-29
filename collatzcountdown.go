package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	i := 0
	for start > 1 {
		if start%2 != 0 {
			start = start*3 + 1
			i++
			continue
		}
		if start%2 == 0 {
			start = start / 2
			i++
			continue
		}
	}
	return i
}
