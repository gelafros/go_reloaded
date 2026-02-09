package function

import (
	"fmt"
	"strconv"
)

// ProcessMarkers handles (hex) and (bin) conversions.
// It iterates through the tokens and applies transformations to the previous token.
func ProcessMarkers(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "(hex)" || token == "(bin)" {
			if i > 0 {
				base := 16
				if token == "(bin)" {
					base = 2
				}

				val, err := strconv.ParseInt(tokens[i-1], base, 64)
				if err == nil {
					tokens[i-1] = fmt.Sprint(val)
					tokens = append(tokens[:i], tokens[i+1:]...)
					i--
				}
			}
		}
	}
	return tokens
}
