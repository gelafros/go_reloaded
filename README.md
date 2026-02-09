# Go Reloaded

A robust, pipeline-driven text formatting tool written in Go.  
This utility automates the cleaning and standardization of raw text files, handling complex casing rules, number base conversions, punctuation logic, and grammatical corrections.

---

## 📖 Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
- [Installation & Usage](#installation--usage)
- [The Rules](#the-rules)
- [Testing](#testing)

---

## Overview

`go-reloaded` is a command-line interface (CLI) tool designed to correct and format text. It parses an input file, applies a series of transformation rules, and outputs the result to a specified file. 

It is built using a **Strict TDD (Test Driven Development)** methodology, ensuring that every feature covers specific edge cases and "Golden Test" scenarios.

## Features

- **Base Conversion**: Automatically detects and converts Hexadecimal and Binary numbers to Decimal.
- **Casing Control**: Apply Uppercase, Lowercase, or Capitalization logic to words or groups of words.
- **Smart Punctuation**: Fixes spacing errors around standard punctuation (`.,!?:;`) and handles grouping characters (`...`, `!?`).
- **Grammar Correction**: Automatically corrects the indefinite article `a` to `an` based on the following word's starting letter.
- **Quote Balancing**: Ensures single quotes are spaced correctly relative to the text they enclose.

## Project Structure

The codebase is organized into modular components to separate the transformation logic (`functions/`), the entry point (`main.go`), and the test suite (`tests/`).

```text
.
├── functions/              # Transformation Logic
│   ├── casing.go           # Handles (up), (low), (cap) logic
│   ├── grammar.go          # Handles vowel checks (a -> an)
│   ├── markers.go          # Handles (hex) and (bin) conversion
│   ├── pipeline.go         # Main Process() function connecting stages
│   ├── punctuation.go      # Handles spacing and quote pairs
│   └── tokenizer.go        # Splits raw text into processable tokens
├── tests/                  # Test Suite
│   ├── casing_test.go
│   ├── grammar_test.go
│   ├── markers_test.go
│   ├── process_test.go     # Integration/Golden tests
│   ├── punctuation_test.go
│   └── tokenizer_test.go
├── go.mod                  # Module definition
├── main.go                 # CLI Entry point (File I/O)
├── README.md               # Documentation
├── sample.txt              # Example input
└── result.txt              # Generated output

Architecture

This project utilizes a Pipeline Architecture rather than a Finite State Machine. This allows for modular, testable stages of processing:

    Tokenizer: Splits raw text into a slice of strings, protecting special groups like ....

    Marker Pass: Scans for transformation triggers (e.g., (hex), (up, 2)) and modifies the token stream.

    Grammar Pass: Analyzes context to fix a vs an.

    Casing Pass: Applies final casing logic.

    Punctuation Pass: Merges punctuation with previous words and handles quote pairs.

    Joiner: Reconstructs the final string.

Installation & Usage
Prerequisites

    Go (version 1.18 or higher recommended)

Running the Tool

Clone the repository and run the program directly using go run:
Bash

# Syntax
go run . <input_file> <output_file>

# Example
go run . sample.txt result.txt

The Rules

The tool scans for specific "marker" strings to trigger formatting.
1. Hex & Binary

Converts the preceding number to decimal.

    1E (hex) -> 30

    10 (bin) -> 2

2. Casing

Modifies the preceding word(s).

    (up): Uppercase (e.g., Ready, set, go (up) -> Ready, set, GO)

    (low): Lowercase (e.g., Stop SHOUTING (low) -> Stop shouting)

    (cap): Capitalize (e.g., welcome (cap) -> Welcome)

Arguments: You can specify a number to affect multiple previous words.

    This is exciting (up, 2) -> This IS EXCITING

3. Punctuation

Standardizes spacing.

    Hello , world ! -> Hello, world!

    I was thinking ... -> I was thinking...

4. Grammar

    a apple -> an apple

    a hour -> an hour (handles 'h')

Testing

The project includes a comprehensive test suite covering unit tests for individual functions and integration tests for the full pipeline.

To run the full suite:
Bash

go test -v ./...

To run the specific "Golden Tests" (Acceptance Criteria):
Bash

go test -v ./tests/ -run TestProcess


---

# Product Requirements Document (PRD) - go-reloaded

## 1. Problem Statement
We need a Go CLI tool that reads an input text file, applies specific formatting/transformation rules (hex/bin conversion, casing, punctuation), and writes the result to an output file.

## 2. User / Use Case
* **Primary User:** A student or auditor running the CLI tool.
* **Use Case:** To normalize and correct raw text files according to a strict set of rules defined by the project subject.

## 3. CLI Contract
* **Command:** `go run . <inputFile> <outputFile>`
* **Inputs:**
    * `<inputFile>`: Path to an existing text file to read.
    * `<outputFile>`: Path where the transformed text will be saved.
* **Output:** A single file containing the processed text.
* **Error Handling:** If arguments are missing or files cannot be read/written, the program should exit with a clear error message.

## 4. Functional Requirements (Rules)

### 4.1 Marker Conversions
* **(hex):** Convert the previous word from Hexadecimal to Decimal.
* **(bin):** Convert the previous word from Binary to Decimal.

### 4.2 Case Transformations
* **(up):** Convert the previous word to Uppercase.
* **(low):** Convert the previous word to Lowercase.
* **(cap):** Capitalize the previous word.
* **(marker, N):** Apply the transformation to the previous N words (e.g., `(up, 2)`).

### 4.3 Punctuation & Quotes
* **Standard Punctuation:** `. , ! ? : ;` must adhere to the previous word and be followed by a space.
* **Groups:** `...` and `!?` must remain grouped.
* **Quotes:** `'` acts as a pair. The first adheres to the next word; the second adheres to the previous word.

### 4.4 Grammar
* **Vowels:** Replace `a` with `an` if the next word begins with a vowel (`a, e, i, o, u`) or `h`.

## 5. Non-Goals
* No usage of external packages (standard library only).
* No complex natural language processing beyond the specified rules.

## 6. Documentation Links
* **[Acceptance Criteria & Golden Tests](./golden_tests.md)** - The specific input/output cases used to verify the rules.
* **[Architecture Decision](./architecture.md)** - The technical approach (Pipeline vs FSM).
* **[Milestone Plan](./milestones.md)** - The implementation roadmap.

## 7. Risks
* **Tokenization:** Splitting text correctly so that punctuation attached to words