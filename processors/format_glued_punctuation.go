package processors

import "strings"

func Format_Glued_Punctuation(s []string) []string {

	for idx, word := range s {

		if idx > 0 && len(word) > 0 {
			if strings.HasPrefix(word, "...") {
				s[idx] = strings.TrimPrefix(s[idx], "...")
				s[idx-1] = s[idx-1] + "..."
			} else if strings.HasPrefix(word, "!?") {
				s[idx] = strings.TrimPrefix(s[idx], "!?")
				s[idx-1] = s[idx-1] + "!?"
			} else {
				first_letter := string(word[0])
				switch first_letter {
				case ",", ":", ";", "!", "?", "/":

					s[idx] = strings.TrimPrefix(s[idx], first_letter)
					s[idx-1] = s[idx-1] + first_letter
				}
			}
		}

	}

	return s
}
