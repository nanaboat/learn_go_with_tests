package iteration

import (
	"bytes"
	"fmt"
	"strings"
)

// Repeat returns a charater repeated a given
// number of times
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

// Repeat returns a charater repeated a given
// number of times
func RepeatString(character string, count int) string {
	return strings.Repeat(character, count)
}

// Repeat returns a charater repeated a given
// number of times
func RepeatString1(character string, count int) string {
	var sb strings.Builder
	for i := 0; i < count; i++ {
		sb.WriteString(character)
	}
	return sb.String()
}

// Repeat returns a charater repeated a given
// number of times
func RepeatString2(character string, count int) string {
	repeated := character
	for i := 0; i < count-1; i++ {
		repeated = fmt.Sprintf("%s%s", repeated, character)
	}
	return repeated
}

// Repeat returns a charater repeated a given
// number of times
func RepeatString3(character string, count int) string {
	var b bytes.Buffer
	for i := 0; i < count; i++ {
		b.WriteString(character)
	}
	return b.String()
}

// Repeat returns a charater repeated a given
// number of times
func RepeatString4(character string, count int) string {
	var s []string
	for i := 0; i < count; i++ {
		s = append(s, character)
	}
	return strings.Join(s, "")
}
