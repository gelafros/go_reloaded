package function

import (
	"strconv"
	"strings"
)

// ProcessCasing handles (up), (low), and (cap) transformations.
// It supports optional numeric arguments (e.g., (up, 2)).
func ProcessCasing(tokens []string) []string {
	result := []string{}
	i := 0
	for i < len(tokens) {
		token := tokens[i]

		// Single markers: (up), (low), (cap)
		if isSingleCasingMarker(token) {
			if len(result) > 0 {
				last := result[len(result)-1]
				if strings.ContainsAny(last, ".,!?:;") && len(last) == 1 {
					result = append(result, token)
					i++
					continue
				}
				applyCasing(result, len(result)-1, len(result)-1, token)
			}
			i++
			continue
		}

		// Markers with arguments: (up, 2)
		// Tokenizer produces: ["(up", ",", "2)"]
		if (token == "(up" || token == "(low" || token == "(cap") && i+2 < len(tokens) {
			if tokens[i+1] == "," && strings.HasSuffix(tokens[i+2], ")") {
				numStr := strings.TrimSuffix(tokens[i+2], ")")
				count, err := strconv.Atoi(numStr)

				if err == nil && count > 0 {
					start := len(result) - count
					if start < 0 {
						start = 0
					}
					end := len(result) - 1
					if end >= start {
						typeStr := token + ")"
						applyCasing(result, start, end, typeStr)
					}
				}
				// Remove 3 tokens: (up , 2)
				i += 3
				continue
			}
		}

		result = append(result, token)
		i++
	}
	return result
}

func isSingleCasingMarker(s string) bool {
	return s == "(up)" || s == "(low)" || s == "(cap)"
}

func applyCasing(tokens []string, start, end int, marker string) {
	for j := start; j <= end; j++ {
		if _, err := strconv.Atoi(tokens[j]); err == nil {
			continue
		}
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
