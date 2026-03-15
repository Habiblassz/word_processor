package processors

func Format_Isolated_Punctuation(s []string) []string {

	for idx, word := range s {

		switch word {
		case "...", ",", "?", ".", "!", ":", ";":
			if idx > 0 {
				s[idx-1] = s[idx-1] + word
				s[idx] = ""
			}
		}

	}

	return s
}
