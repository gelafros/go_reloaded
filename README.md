# go-reloaded

A production-grade text formatting engine built with Go, implementing a pipeline-driven architecture for automated text normalization and transformation.

[![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go)](https://go.dev/)
[![Test Coverage](https://img.shields.io/badge/tests-50%20cases-success)](./tests/)
[![Architecture](https://img.shields.io/badge/architecture-pipeline-blue)](./docs/architecture.md)

---

## Overview

**go-reloaded** is a command-line text processing utility that applies sophisticated transformation rules to raw text files. Built using strict Test-Driven Development (TDD) methodology, it handles number base conversions, intelligent casing transformations, punctuation normalization, and grammatical corrections through a modular pipeline architecture.

### Key Capabilities

- **Number Base Conversion**: Automatic hexadecimal and binary to decimal conversion
- **Intelligent Casing**: Context-aware uppercase, lowercase, and capitalization with multi-word support
- **Smart Punctuation**: Automated spacing correction and punctuation grouping (`...`, `!?`)
- **Grammar Correction**: Automatic indefinite article correction (`a` → `an`) based on phonetic rules
- **Quote Balancing**: Intelligent single-quote pairing and spacing

---

## Quick Start

### Prerequisites

- Go 1.22 or higher
- Unix-like environment (Linux, macOS) or Windows with Go installed

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd go-reloaded

# Verify installation
go version
```

### Usage

```bash
# Basic syntax
go run . <input_file> <output_file>

# Example
go run . sample.txt result.txt
```

### Example Transformation

**Input** (`sample.txt`):
```text
1E (hex) files were added
Ready, set, go (up) !
There it was. A amazing rock!
```

**Output** (`result.txt`):
```text
30 files were added
Ready, set, GO!
There it was. An amazing rock!
```

---

## Transformation Rules

### 1. Number Base Conversion

Converts numbers from hexadecimal or binary to decimal format.

| Marker | Description | Example |
|--------|-------------|---------|
| `(hex)` | Hexadecimal to decimal | `1E (hex)` → `30` |
| `(bin)` | Binary to decimal | `10 (bin)` → `2` |

**Behavior**: The marker applies to the immediately preceding token and is removed from output.

### 2. Case Transformations

Modifies the case of preceding words with optional multi-word support.

| Marker | Description | Example |
|--------|-------------|---------|
| `(up)` | Convert to UPPERCASE | `hello (up)` → `HELLO` |
| `(low)` | Convert to lowercase | `WORLD (low)` → `world` |
| `(cap)` | Capitalize first letter | `john (cap)` → `John` |
| `(marker, N)` | Apply to N previous words | `hello world (up, 2)` → `HELLO WORLD` |

**Multi-word syntax**: `(up, 3)` applies uppercase to the previous 3 words.

### 3. Punctuation Normalization

Automatically corrects spacing around punctuation marks.

**Standard Punctuation**: `. , ! ? : ;`
- Adheres to the previous word (no space before)
- Followed by a space (or end of text)

**Punctuation Groups**: `...`, `!?`, `?!`
- Treated as single units
- Follow standard punctuation rules

**Examples**:
```text
Hello , world !     →  Hello, world!
I was thinking ...  →  I was thinking...
Really ?!           →  Really?!
```

### 4. Quote Handling

Single quotes (`'`) are treated as balanced pairs with intelligent spacing.

**Rules**:
- Opening quote: Adheres to the following word
- Closing quote: Adheres to the preceding word
- Quotes alternate between opening and closing

**Examples**:
```text
' hello '                    →  'hello'
I said ' hello ' to ' you '  →  I said 'hello' to 'you'
```

### 5. Grammar Correction

Automatically corrects indefinite articles based on phonetic rules.

**Rule**: `a` → `an` when the following word starts with:
- Vowels: `a, e, i, o, u`
- The letter `h` (silent h rule)

**Examples**:
```text
a apple   →  an apple
a hour    →  an hour
a book    →  a book
```

**Case preservation**: `A apple` → `An apple`

---

## Architecture

### Design Philosophy

**go-reloaded** implements a **Pipeline Architecture** rather than a Finite State Machine (FSM), chosen for its superior modularity, testability, and ability to handle look-behind operations efficiently.

### Pipeline Stages

```text
┌─────────────┐
│ Input File  │
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────────────┐
│ 1. Tokenizer                            │
│    Splits text into processable tokens  │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│ 2. Marker Processing                    │
│    Handles (hex), (bin) conversions     │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│ 3. Grammar Processing                   │
│    Applies a → an corrections           │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│ 4. Casing Processing                    │
│    Applies (up), (low), (cap) rules     │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│ 5. Punctuation Processing               │
│    Normalizes spacing and quotes        │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│ 6. Joiner                               │
│    Reconstructs final string            │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────┐
│ Output File │
└─────────────┘
```

### Why Pipeline Over FSM?

| Aspect | Pipeline | FSM |
|--------|----------|-----|
| **Complexity** | Low - isolated stages | High - complex state transitions |
| **Look-behind** | Easy - direct array access | Hard - requires manual buffering |
| **Testing** | Easy - unit test each stage | Medium - requires full sequence testing |
| **Maintainability** | High - clear separation of concerns | Medium - tightly coupled states |

---

## Project Structure

```text
go-reloaded/
├── functions/              # Core transformation logic
│   ├── pipeline.go         # Main orchestration (Process function)
│   ├── tokenizer.go        # Text → []string conversion
│   ├── markers.go          # (hex), (bin) conversion logic
│   ├── grammar.go          # a → an correction logic
│   ├── casing.go           # (up), (low), (cap) logic
│   └── punctuation.go      # Spacing and quote handling
│
├── tests/                  # Comprehensive test suite
│   ├── integration_test.go # 50 golden test cases
│   ├── tokenizer_test.go   # Tokenization unit tests
│   ├── markers_test.go     # Conversion unit tests
│   ├── grammar_test.go     # Grammar unit tests
│   ├── casing_test.go      # Casing unit tests
│   └── punctuation_test.go # Punctuation unit tests
│
├── docs/                   # Technical documentation
│   ├── architecture.md     # Architecture decisions
│   ├── golden_tests.md     # 50 acceptance criteria
│   └── milestones.md       # Development roadmap
│
├── main.go                 # CLI entry point
├── go.mod                  # Go module definition
├── LICENSE.md              # MIT License
├── sample.txt              # Example input file
└── README.md               # This file
```

---

## Testing

### Test-Driven Development (TDD)

This project follows **strict TDD methodology**:
- Every feature has a failing test written first
- Code is written to pass the test
- Refactoring maintains green tests

### Test Coverage

The test suite includes **50 comprehensive test cases** covering:

1. **Audit Cases** (11 tests): Core functionality from project requirements
2. **Standard Edge Cases** (8 tests): Common variations
3. **Stress Tests** (8 tests): Boundary conditions and error handling
4. **Clever Combinations** (6 tests): Multi-rule interactions
5. **The Gauntlet** (17 tests): Complex chained transformations

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run integration tests only
go test -v ./tests/ -run TestProcess

# Run specific test suite
go test -v ./tests/ -run TestTokenize
go test -v ./tests/ -run TestMarkers
go test -v ./tests/ -run TestGrammar
go test -v ./tests/ -run TestCasing
go test -v ./tests/ -run TestPunctuation
```

### Example Test Cases

```go
// Simple transformations
"1E (hex) files"              → "30 files"
"Ready, set, go (up) !"       → "Ready, set, GO!"
"A amazing rock!"             → "An amazing rock!"

// Complex chained transformations
"one (up) (low) (cap)"        → "One"
"a (up) apple"                → "AN apple"
"hello world (up, 2) (low, 1)" → "HELLO world"

// Edge cases
"(up) starts with marker"     → "starts with marker"
"hello (up, 100) world"       → "HELLO world"
"' ' ' '"                     → "'' ''"
```

---

## Implementation Details

### Tokenization Strategy

The tokenizer uses a **Protect-Pad-Split** approach:

1. **Protect**: Replace special groups with placeholders
   - `...` → `__ELLIPSIS__`
   - `!?` → `__BANG_QMARK__`

2. **Pad**: Add spaces around standard punctuation
   - `.` → ` . `
   - `,` → ` , `

3. **Restore**: Replace placeholders with original groups

4. **Split**: Use `strings.Fields()` for whitespace splitting

### Marker Processing

Markers are processed in order of appearance:
- Markers modify preceding tokens
- Invalid markers are removed
- Numeric arguments are parsed: `(up, 2)`
- Markers on punctuation are ignored

### Quote Balancing

Quotes maintain state across the token stream:
- First `'` → opening (adheres right)
- Second `'` → closing (adheres left)
- State alternates for each quote encountered

---

## Error Handling

The application handles errors gracefully:

```bash
# Missing arguments
$ go run .
usage: go run . <input_file> <output_file>

# File not found
$ go run . missing.txt output.txt
error reading input file: open missing.txt: no such file or directory

# Permission denied
$ go run . input.txt /root/output.txt
error writing output file: open /root/output.txt: permission denied
```

---

## Development

### Adding New Features

1. **Write the test** in appropriate `tests/*_test.go` file
2. **Run tests** to confirm failure: `go test -v ./tests/`
3. **Implement** the feature in `functions/`
4. **Verify** tests pass
5. **Refactor** while maintaining green tests

### Code Style

- Standard Go formatting (`gofmt`)
- Clear function names describing behavior
- Minimal comments (self-documenting code)
- No external dependencies (stdlib only)

### Contributing

This is an educational project following specific requirements. Modifications should:
- Maintain TDD discipline
- Pass all 50 golden tests
- Use only Go standard library
- Follow the pipeline architecture

---

## Performance Characteristics

- **Time Complexity**: O(n) where n is the number of tokens
- **Space Complexity**: O(n) for token storage
- **Processing**: Single-pass through pipeline stages
- **Memory**: Efficient string slice operations

---

## Limitations

- **Single quotes only**: Double quotes are treated as regular text
- **English-centric**: Grammar rules assume English phonetics
- **No nested markers**: `(up (low))` is not supported
- **Sequential processing**: No parallel pipeline execution

---

## License

MIT License - See [LICENSE.md](./LICENSE.md) for details.

---

## Acknowledgments

- Built using strict TDD methodology
- Architecture inspired by Unix pipeline philosophy
- Test cases derived from project acceptance criteria

---

## Documentation

- **[Architecture Decision](./docs/architecture.md)** - Technical design rationale
- **[Golden Tests](./docs/golden_tests.md)** - Complete test specification (50 cases)
- **[Milestones](./docs/milestones.md)** - Development roadmap
- **[PRD](./PRD.md)** - Product Requirements Document

---

## Support

For issues or questions:
1. Review the [golden tests](./docs/golden_tests.md) for expected behavior
2. Check the [architecture documentation](./docs/architecture.md)
3. Run the test suite to identify specific failures

---

**Built with Go • Tested with 50 Cases • Pipeline Architecture**
