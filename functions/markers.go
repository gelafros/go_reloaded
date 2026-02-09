package functions

import (
	"fmt"
	"strconv"
)

// ProcessMarkers handles (hex) and (bin) conversions.
// It iterates through the tokens and applies transformations to the previous token.
func ProcessMarkers(tokens []string) []string {
	var result []string
	for _, token := range tokens {
		if token == "(hex)" || token == "(bin)" {
			success := false
			if len(result) > 0 {
				base := 16
				if token == "(bin)" {
					base = 2
				}

				val, err := strconv.ParseInt(result[len(result)-1], base, 64)
				if err == nil {
					result[len(result)-1] = fmt.Sprint(val)
					success = true
				}
			}
			if !success {
				result = append(result, "")
			}
		} else {
			result = append(result, token)
		}
	}
	return result
}
