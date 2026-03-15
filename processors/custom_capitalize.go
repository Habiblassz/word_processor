package processors

import "strings"

func Custom_Capitalize(s string) string {

	if len(s) == 0 {
		return s
	}

	first_letter := strings.ToUpper(string(s[0]))
	remaining_letters := strings.ToLower(s[1:])

	return first_letter + remaining_letters
}
