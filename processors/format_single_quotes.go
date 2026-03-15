package processors

import "strings"

func Format_Single_Quotes(s []string) []string {

	is_quote_open := false

	for idx := 0; idx < len(s); idx++ {
		if s[idx] == "'" && idx+1 < len(s) && is_quote_open == false {
			s[idx] = ""
			s[idx+1] = "'" + s[idx+1]
			is_quote_open = true
			idx++
		} else if s[idx] == "'" && idx > 0 && is_quote_open {
			s[idx] = ""
			s[idx-1] = s[idx-1] + "'"
			is_quote_open = false
		} else if idx > 0 && is_quote_open == false {
			if strings.HasPrefix(s[idx], "'") {
				is_quote_open = true
			}
		} else if idx > 0 && is_quote_open {
			if strings.HasPrefix(s[idx], "'") {
				s[idx] = strings.TrimPrefix(s[idx], "'")
				s[idx-1] = s[idx-1] + "'"
				is_quote_open = false
			}
		}
	}
	return s
}
