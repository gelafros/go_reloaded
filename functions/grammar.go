package function

import "strings"

// ProcessGrammar handles the 'a' -> 'an' rule.
func ProcessGrammar(tokens []string) []string {
	for i, token := range tokens {
		if token == "a" || token == "A" {
			// Look ahead for the next word, skipping markers
			nextWord := ""
			for j := i + 1; j < len(tokens); j++ {
				if !isMarkerGrammar(tokens[j]) {
					nextWord = tokens[j]
					break
				}
			}

			if nextWord != "" {
				firstChar := rune(nextWord[0])
				if isVowel(firstChar) {
					if token == "a" {
						tokens[i] = "an"
					} else {
						tokens[i] = "An"
					}
				}
			}
		}
	}
	return tokens
}

func isMarkerGrammar(s string) bool {
	return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouhAEIOUH", r)
}
