package processors

func Format_Article(s []string) []string {

	for idx, word := range s {
		switch word {
		case "A", "a":

			if idx+1 < len(s) && len(s[idx+1]) > 0 {
				first_letter := string(s[idx+1][0])

				switch first_letter {
				case "a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H":
					if word == "a" {
						s[idx] = "an"
					} else {
						s[idx] = "An"
					}
				}
			}
		}
	}

	return s
}
