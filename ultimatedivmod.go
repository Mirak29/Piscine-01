package piscine

func UltimateDivMod(a *int, b *int) {
	i := *a
	*a = *a / *b
	*b = i % *b
}
