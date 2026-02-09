# Go-Reloaded

**Go-Reloaded** is a robust text formatting tool built in Go. It automates the process of correcting, formatting, and transforming text files based on a predefined set of rules and embedded markers.

## 🚀 Features

### 1. Numeric Conversions
Convert hexadecimal and binary values to decimal automatically.
*   **(hex)**: Converts the preceding hexadecimal number to decimal.
    *   `1E (hex)` → `30`
*   **(bin)**: Converts the preceding binary number to decimal.
    *   `10 (bin)` → `2`

### 2. Casing Transformations
Modify the case of words using embedded markers. Supports applying changes to multiple preceding words.
*   **(up)**: Uppercase the previous word.
    *   `go (up)` → `GO`
*   **(low)**: Lowercase the previous word.
    *   `SHOUTING (low)` → `shouting`
*   **(cap)**: Capitalize the previous word.
    *   `bridge (cap)` → `Bridge`
*   **Arguments**: `(up, 2)` transforms the previous 2 words.

### 3. Punctuation & Grammar
Ensure professional formatting and correct grammar.
*   **Punctuation**: Fixes spacing for `. , ! ? : ;`.
    *   `hello , world` → `hello, world`
*   **Groups**: Preserves punctuation groups like `...` and `!?`.
*   **Quotes**: Correctly spaces single quotes `'` around words.
    *   `' word '` → `'word'`
*   **Vowels**: Corrects `a` to `an` before vowels or `h`.
    *   `a apple` → `an apple`

## 📦 Installation & Usage

### Prerequisites
*   Go installed on your machine.

### Running the Tool
Clone the repository and run the program directly using `go run`.

```bash
go run . <input_file> <output_file>
```

## Features
* Hex/Bin conversion
* Casing transformations (up, low, cap)
* Punctuation formatting
* Grammar correction (a -> an)

## Examples -Here are some examples of how the tool transforms text: +* input_file: Path to the file containing text to process. +* output_file: Path where the result will be written.

-| Input | Output | Rule | -|-------|--------|------| -| 1E (hex) files were added | 30 files were added | Hex Conversion | -| Ready, set, go (up) ! | Ready, set, GO! | Uppercase | -| I was thinking ... You were right | I was thinking... You were right | Punctuation Groups | -| There it was. A amazing rock! | There it was. An amazing rock! | Grammar (a -> an) | +### Example

-## Features -* Hex/Bin conversion: Converts numbers preceding (hex) or (bin) to decimal. -* Casing transformations: (up), (low), (cap) modify the preceding word(s). -* Punctuation formatting: Fixes spacing around punctuation and handles quotes. -* Grammar correction: Automatically corrects 'a' to 'an' before vowels. +Command: +bash +go run . sample.txt result.txt + + +Input (sample.txt): +text +1E (hex) files were added. +It has been 10 (bin) years. +Ready, set, go (up) ! +I was thinking ... You were right +There it was. A amazing rock! + + +Output (result.txt): +text +30 files were added. +It has been 2 years. +Ready, set, GO! +I was thinking... You were right +There it was. An amazing rock! + + +## 🧪 Testing + +The project includes a comprehensive test suite covering unit logic and integration scenarios. + +Run all tests with: +bash +go test -v ./... + + +## 📂 Architecture + +The application follows a Pipeline architecture, processing text in distinct stages (Tokenization → Markers → Grammar → Casing → Punctuation) to ensure deterministic results and easy testing.