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

## High-Level Design (The Pipeline)

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
```

# Tokenization Specification (Pad & Split)

## Protect Groups
Identify protected groups such as `...` and `!?` first.

## Pad Punctuation
Replace punctuation characters (e.g. `.`, `,`, `!`, etc.) with a padded form:


## Split
Use `strings.Fields()` to create the final token slice.

---

# Implementation Methodology (Strict TDD)

This project strictly follows **Test-Driven Development (TDD)**.  
No feature code is allowed to exist without a preceding failing test.

## Development Cycle

1. **Select Task**  
   Open the relevant task file:  

2. **Read Script**  
   Follow the **TDD Execution Script** inside the task card.

3. **Red**  
   Write the unit test using inputs from:  
   `docs/golden_tests.md`

4. **Green**  
   Write the *minimum* amount of code required to pass the test.

5. **Refactor**  
   Clean up the implementation while keeping tests green.

---

## Enforcement

Violation of this protocol (writing code before tests) is considered an **architectural failure**.
