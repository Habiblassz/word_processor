# Go Text Formatter (Go Reloaded)

A lightweight, modular text processing and auto-correction tool written in Go. This program takes an input text file, applies a strict set of formatting rules and grammatical corrections, and writes the polished result to an output file.

## Features

This tool automatically detects and applies the following formatting rules:

### 1. Grammatical Corrections

* **Vowel Matching:** Automatically converts the article `a` (or `A`) to `an` (or `An`) if the following word begins with a vowel or the letter 'h'.

### 2. Base Conversions

* `(hex)`: Replaces the preceding word with its decimal equivalent (parsing it as a hexadecimal number).
* `(bin)`: Replaces the preceding word with its decimal equivalent (parsing it as a binary number).

### 3. Text Case Manipulation

* `(up)`: Converts the preceding word to UPPERCASE.
* `(low)`: Converts the preceding word to lowercase.
* `(cap)`: Capitalizes the first letter of the preceding word.
* **Multi-word manipulation:** Adding a number inside the tag applies the formatting to the specified number of preceding words (e.g., `(up, 2)`, `(low, 3)`, `(cap, 6)`).

### 4. Punctuation Formatting

* **Standard Punctuation:** Ensures that marks like `.`, `,`, `!`, `?`, `:`, and `;` are attached directly to the word preceding them, followed by a space.
* **Grouped Punctuation:** Preserves grouped punctuation like `...` or `!?` and formats them as a single block attached to the previous word.
* **Single Quotes (`'`):** Acts as a toggle switch. It attaches the first quote to the beginning of the next word, and the second quote to the end of the previous word (e.g., `' hello '` becomes `'hello'`).

## Project Structure

The project is built using a modular architecture to ensure separation of concerns, readability and testability:

* **Manager Function:** Handles file I/O and passes data through a sequential pipeline.
* **Worker Functions:** Isolated functions handling specific tasks (e.g., `format_tags`, `format_isolated_punctuation`, `format_single_quotes`).

## Installation & Usage

1. **Clone the repository:**
```bash
git clone https://github.com/Habiblassz/word_processor.git
cd word_processor

```


2. **Build the executable:**
```bash
go build -o text_parser .

```


3. **Run the program:**
Provide the name of the input file and the desired output file as arguments.
```bash
./text_parser sample.txt result.txt

```


*(Alternatively, you can run it directly without building: `go run . sample.txt result.txt`)*
