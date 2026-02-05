# Architecture Decision

## Decision Summary
* **Selected Architecture:** **Pipeline**
* **Data Structure:** `[]string` (Slice of strings)

## Rationale
We chose the Pipeline approach over a Finite State Machine (FSM). While FSMs are efficient for streaming data, the Pipeline model fits the specific requirements of this project (modifying "previous" words) much better.

### Comparison: Pipeline vs FSM

| Feature | Pipeline (`[]string`) | Finite State Machine (FSM) |
| :--- | :--- | :--- |
| **Complexity** | **Low.** Logic is split into isolated stages. | **High.** Requires tracking complex state transitions. |
| **Look-Behind** | **Easy.** Can verify `words[i-1]` instantly. | **Hard.** Requires buffering previous tokens manually. |
| **Testing** | **Easy.** Each stage can be unit tested separately. | **Medium.** Must test the entire stream sequence. |
| **Performance** | Medium. Requires multiple passes over the data. | High. Single pass (O(n)). |
| **Memory** | Medium. Loads entire file into memory. | Low. Processes stream byte-by-byte. |

### Why Pipeline Wins
1.  **Simplicity:** The logic primarily involves looking at a word and modifying the *previous* word (e.g., for `(hex)` or `(up)`). A slice of strings allows easy index lookups.
2.  **Testability:** We can break the processing into distinct functions (e.g., `ProcessMarkers`, `ProcessPunctuation`) and test them independently with TDD.
3.  **Context:** The performance cost of multiple passes is negligible for the expected text file sizes in this project.

## High-Level Design (Sketch)

The program will flow through these stages:

```text
[Input File] 
    |
    v
[Tokenizer] -> Converts text into []string (separating punctuation)
    |
    v
[Pass 1: Markers] -> Handles (hex), (bin), (up), (low), (cap)
    |
    v
[Pass 2: Punctuation & Quotes] -> Fixes spacing around symbols
    |
    v
[Pass 3: Grammar] -> Fixes 'a' vs 'an'
    |
    v
[Joiner] -> Combines []string into single string
    |
    v
[Output File]