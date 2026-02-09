package function

import (
	"strconv"
	"strings"
)

// ProcessCasing handles (up), (low), and (cap) transformations.
// It supports optional numeric arguments (e.g., (up, 2)).
func ProcessCasing(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Single markers: (up), (low), (cap)
		if token == "(up)" || token == "(low)" || token == "(cap)" {
			if i > 0 {
				// Check if previous token is a marker (orphan protection)
				if isMarker(tokens[i-1]) {
					continue
				}
				// Golden Test #33: Marker broken by punctuation (ignored/preserved)
				if strings.ContainsAny(tokens[i-1], ".,!?:;") && len(tokens[i-1]) == 1 {
					continue
				}
				applyCasing(tokens, i-1, i-1, token)
				// Remove marker
				tokens = append(tokens[:i], tokens[i+1:]...)
				i--
			}
			continue
		}

		// Markers with arguments: (up, 2)
		// Tokenizer produces: ["(up", ",", "2)"]
		if (token == "(up" || token == "(low" || token == "(cap") && i+2 < len(tokens) {
			if tokens[i+1] == "," && strings.HasSuffix(tokens[i+2], ")") {
				numStr := strings.TrimSuffix(tokens[i+2], ")")
				count, err := strconv.Atoi(numStr)

				if err == nil && count >= 0 {
					// Valid marker syntax
					start := i - count
					if start < 0 {
						start = 0
					}
					end := i - 1

					// Check for markers in range to prevent breaking chains
					hasMarker := false
					for k := start; k <= end; k++ {
						if isMarker(tokens[k]) {
							hasMarker = true
							break
						}
					}

					if !hasMarker {
						if count > 0 {
							typeStr := token + ")"
							applyCasing(tokens, start, end, typeStr)
						}
						// Remove 3 tokens: (up , 2)
						tokens = append(tokens[:i], tokens[i+3:]...)
						i--
					}
				}
			}
		}
	}
	return tokens
}

func isMarker(s string) bool {
	return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}

func applyCasing(tokens []string, start, end int, marker string) {
	for j := start; j <= end; j++ {
		switch marker {
		case "(up)":
			tokens[j] = strings.ToUpper(tokens[j])
		case "(low)":
			tokens[j] = strings.ToLower(tokens[j])
		case "(cap)":
			if len(tokens[j]) > 0 {
				tokens[j] = strings.ToUpper(string(tokens[j][0])) + strings.ToLower(tokens[j][1:])
			}
		}
	}
}
