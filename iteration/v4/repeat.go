package iteration

// Repeat returns a charater repeated a given
// number of times
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
