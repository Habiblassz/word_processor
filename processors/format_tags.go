package processors

import (
	"fmt"
	"strconv"
	"strings"
)

func Format_Tags(s []string) []string {

	for idx, word := range s {
		switch word {
		case "(hex)":
			if idx > 0 {

				decimal_integer, decimal_error := strconv.ParseInt(s[idx-1], 16, 64)

				if decimal_error != nil {
					fmt.Println("Error. Unable to convert to decimal integer")
				}

				decimal_string := strconv.FormatInt(decimal_integer, 10)
				s[idx-1] = decimal_string
				s[idx] = ""
			}

		case "(bin)":
			if idx > 0 {

				decimal_integer, decimal_error := strconv.ParseInt(s[idx-1], 2, 64)

				if decimal_error != nil {
					fmt.Println("Error. Unable to convert to decimal integer")
				}

				decimal_string := strconv.FormatInt(decimal_integer, 10)
				s[idx-1] = decimal_string
				s[idx] = ""
			}

		case "(low)":
			if idx > 0 {
				s[idx-1] = strings.ToLower(s[idx-1])
				s[idx] = ""
			}
		case "(low,":
			if idx+1 < len(s) {
				iterator, iterator_error := strconv.Atoi(strings.TrimSuffix(s[idx+1], ")"))

				if iterator_error != nil {
					fmt.Println("Unable to convert iterator to integer")
				}

				for i := 1; i <= iterator; i++ {
					if idx-i >= 0 {
						s[idx-i] = strings.ToLower(s[idx-i])
					}
				}

				s[idx] = ""
				s[idx+1] = ""

			}
		case "(up)":
			if idx > 0 {
				s[idx-1] = strings.ToUpper(s[idx-1])
				s[idx] = ""
			}
		case "(up,":
			if idx+1 < len(s) {
				iterator, iterator_error := strconv.Atoi(strings.TrimSuffix(s[idx+1], ")"))

				if iterator_error != nil {
					fmt.Println("Unable to convert iterator to integer")
				}

				for i := 1; i <= iterator; i++ {
					if idx-i >= 0 {
						s[idx-i] = strings.ToUpper(s[idx-i])
					}
				}

				s[idx] = ""
				s[idx+1] = ""

			}
		case "(cap)":
			if idx > 0 {
				s[idx-1] = Custom_Capitalize(s[idx-1])
				s[idx] = ""
			}
		case "(cap,":
			if idx+1 < len(s) {
				iterator, iterator_error := strconv.Atoi(strings.TrimSuffix(s[idx+1], ")"))

				if iterator_error != nil {
					fmt.Println("Unable to convert iterator to integer")
				}

				for i := 1; i <= iterator; i++ {
					if idx-i >= 0 {
						s[idx-i] = Custom_Capitalize(s[idx-i])
					}
				}

				s[idx] = ""
				s[idx+1] = ""

			}
		}
	}

	return s
}
