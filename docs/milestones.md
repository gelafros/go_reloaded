# Milestone Plan

This plan breaks the implementation into testable steps.

## Milestone 1: Setup & Tokenization
* **Goal:** Create the project structure and read files.
* **Deliverable:** A program that reads input, splits it into tokens (words/punctuation), and writes it back to output unchanged.
* **Test:** Input `Hello, world!` -> Output `Hello , world !` (or similar tokenized view).

## Milestone 2: Marker Logic
* **Goal:** Implement the logic for `(hex)`, `(bin)`, `(up)`, `(low)`, and `(cap)`.
* **Deliverable:** Program passes Golden Tests #1 through #6, #12-16.

## Milestone 3: Formatting (Punctuation & Quotes)
* **Goal:** Implement the logic for spacing around punctuation and handling quote pairs.
* **Deliverable:** Program passes Golden Tests #7 through #10, #17-18.

## Milestone 4: Grammar Logic
* **Goal:** Implement `a` -> `an` replacement.
* **Deliverable:** Program passes Golden Test #11 and #19.

## Milestone 5: Final Integration & Audit
* **Goal:** Run all Golden Tests and verify against the provided `sample.txt`.
* **Deliverable:** Final executable ready for peer audit.