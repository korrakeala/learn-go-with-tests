package iteration

import "strings"

func Repeat(character string, iterations int) string {
	/*var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated*/
	return strings.Repeat(character, iterations)
}
