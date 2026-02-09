package function

// ProcessPunctuation handles spacing around punctuation and quotes.
func ProcessPunctuation(tokens []string) []string {
	punct := map[string]bool{
		".": true, ",": true, "!": true, "?": true, ":": true, ";": true,
		"...": true, "?!": true, "!?": true,
	}

	openQuote := false

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Handle Quotes
		if token == "'" {
			if openQuote {
				// Closing quote: adheres to previous
				if i > 0 {
					tokens[i-1] += "'"
					tokens = append(tokens[:i], tokens[i+1:]...)
					i--
				}
				openQuote = false
			} else {
				// Opening quote: adheres to next
				// Special case: ' ' -> '' (Empty quotes pair)
				if i+1 < len(tokens) && tokens[i+1] == "'" {
					tokens[i+1] = "''"
					tokens = append(tokens[:i], tokens[i+1:]...)
					i--
				} else if i+1 < len(tokens) {
					tokens[i+1] = "'" + tokens[i+1]
					tokens = append(tokens[:i], tokens[i+1:]...)
					i--
					openQuote = true
				}
			}
			continue
		}

		// Handle Standard Punctuation
		if punct[token] {
			if i > 0 {
				tokens[i-1] += token
				tokens = append(tokens[:i], tokens[i+1:]...)
				i--
			}
		}
	}
	return tokens
}
