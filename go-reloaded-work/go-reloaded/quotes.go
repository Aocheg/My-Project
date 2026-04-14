package main

func HandleQuotes(words []string) []string {
	var result []string
	inQuote := false

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "'" {
			if inQuote {
				if len(result) > 0 {
					result[len(result)-1] = result[len(result)-1] + "'"
				} else {
					result = append(result, "'")
				}
			} else {
				result = append(result, "'")
			}
			inQuote = !inQuote
			continue
		}
		if inQuote {
			if len(result) > 0 {
				result[len(result)-1] = result[len(result)-1] + word
			} else {
				result = append(result, word)
			}
		} else {
			result = append(result, word)
		}
	}
	return result
}
