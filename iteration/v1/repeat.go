package iteration

// Repeat returns a charater repeated 5 times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
