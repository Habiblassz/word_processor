package processors

func Remove_Space(s []string) []string {
	var result []string
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}

	return result
}
