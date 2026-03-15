package main

import (
	"fmt"
	"os"
	"strings"
	"word_processor/processors"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Incomplete arguments. Add input and output files")
		return
	}

	input_file := os.Args[1]
	output_file := os.Args[2]

	byte_input_file, content_error := os.ReadFile(input_file)

	if content_error != nil {
		fmt.Println("Error. Unable to read the content of", input_file)
		return
	}

	string_input_file := string(byte_input_file)

	words := strings.Fields(string_input_file)
	words = processors.Format_Tags(words)
	words = processors.Remove_Space(words)
	words = processors.Format_Article(words)
	words = processors.Format_Isolated_Punctuation(words)
	words = processors.Remove_Space(words)
	words = processors.Format_Glued_Punctuation(words)
	words = processors.Format_Single_Quotes(words)
	words = processors.Remove_Space(words)

	formatted_words := strings.Join(words, " ")

	write_error := os.WriteFile(output_file, []byte(formatted_words), 0644)

	if write_error == nil {
		fmt.Println("Successfully formatted", input_file, "and printed result in", output_file)
	}

}
