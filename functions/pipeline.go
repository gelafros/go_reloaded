package function

import "strings"

// Process takes the raw input string and runs it through the full pipeline.
func Process(input string) string {
	tokens := Tokenize(input)
	tokens = ProcessMarkers(tokens)
	tokens = ProcessGrammar(tokens)
	tokens = ProcessCasing(tokens)
	tokens = ProcessPunctuation(tokens)
	return strings.Join(tokens, " ")
}
