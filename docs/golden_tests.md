# Golden Tests (Acceptance Criteria)

These tests define the "Done" state. Every test case must pass for the project to be successful.
**Truth Rule:** If the code output disagrees with this file, the code is wrong.

## 1. Audit Cases (Required)
These are taken directly from the project subject.

| ID | Input | Expected Output | Rule |
|----|-------|-----------------|------|
| 1 | `1E (hex) files were added` | `30 files were added` | Hex Conversion |
| 2 | `It has been 10 (bin) years` | `It has been 2 years` | Bin Conversion |
| 3 | `Ready, set, go (up) !` | `Ready, set, GO!` | Uppercase |
| 4 | `I should stop SHOUTING (low)` | `I should stop shouting` | Lowercase |
| 5 | `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` | Capitalization |
| 6 | `This is so exciting (up, 2)` | `This is SO EXCITING` | Multi-word Casing |
| 7 | `I was sitting over there ,and then BAMM !!` | `I was sitting over there, and then BAMM!!` | Punctuation Spacing |
| 8 | `I was thinking ... You were right` | `I was thinking... You were right` | Punctuation Groups |
| 9 | `I am exactly how they describe me: ' awesome '` | `I am exactly how they describe me: 'awesome'` | Quotes |
| 10 | `As Elton John said: ' I am the most well-known homosexual in the world '` | `As Elton John said: 'I am the most well-known homosexual in the world'` | Long Quotes |
| 11 | `There it was. A amazing rock!` | `There it was. An amazing rock!` | Vowels (a -> an) |

## 2. Standard Edge Cases
Common variations of the rules.

| ID | Input | Expected Output | Rule |
|----|-------|-----------------|------|
| 12 | `I bought 1010 (hex) apples` | `I bought 4112 apples` | Large Hex Number |
| 13 | `I have known him for 10100 (bin) years` | `I have known him for 20 years` | Large Bin Number |
| 14 | `Please make this loud (up)` | `Please make this LOUD` | Standard Up |
| 15 | `This is a MAJOR (low) mistake` | `This is a major mistake` | Standard Low |
| 16 | `my name is john doe (cap, 2)` | `my name is John Doe` | Multi-word Cap |
| 17 | `Are you sure ???` | `Are you sure???` | Triple Punctuation |
| 18 | `This is ' tricky '` | `This is 'tricky'` | Simple Quote |
| 19 | `It has a high value` | `It has an high value` | Vowels (h case) |

## 3. Stress Tests (Stability Checks)
Tests for index out of bounds, invalid markers, and empty inputs.

| ID | Input | Expected Output | Rule |
|----|-------|-----------------|------|
| 20 | `(up) starts with marker` | `starts with marker` | Marker at start (Removed) |
| 21 | `Hello (up, 5) world` | `HELLO world` | Count > Words Available (Safe Cap) |
| 22 | `(hex) 1E` | `1E` | Marker before number (Removed) |
| 23 | `User (cap, 0) test` | `User test` | Count is zero (Marker Removed) |
| 24 | `Complex ... ?! strings` | `Complex...?! strings` | Mixed Punctuation Groups |
| 25 | `Multiple ' quotes ' in ' one ' line` | `Multiple 'quotes' in 'one' line` | Multiple Quote Pairs |
| 26 | `Nested (hex) (up)` | `Nested` | Broken/Orphaned Markers (Removed) |
| 27 | `  Spaces   everywhere  ,  right?  ` | `Spaces everywhere, right?` | Excess Whitespace Cleanup |

## 4. Clever Combinations (The "Logic" Trap)
Tests that require multiple rules to fire in the correct order.

| ID | Input | Expected Output | Rule |
|----|-------|-----------------|------|
| 28 | `1e (hex) (up) is thirty` | `30 is thirty` | Hex then Up (Chain Reaction) |
| 29 | `it (cap) (up) worked` | `IT worked` | Cap then Up (Last wins or Cumulative) |
| 30 | `A (up) amazing result` | `AN amazing result` | Up, THEN Vowel Check (A->An) |
| 31 | `He said: ' 10 (bin) '` | `He said: '2'` | Quote + Binary |
| 32 | `... (up)` | `...` | Marker on Punctuation (Ignored) |
| 33 | `word , (up)` | `word, (up)` | Marker broken by punctuation |

## 5. The Gauntlet (Complex Interactions)
These tests ensure robustness against chained, overlapping, or chaotic inputs.

| ID | Input | Expected Output | Rule |
|----|-------|-----------------|------|
| 34 | `one (up) (low) (cap)` | `One` | Chain: ONE -> one -> One |
| 35 | `a (up) apple` | `AN apple` | Up makes "A", Vowel makes "An", preserve case |
| 36 | `a (low) apple` | `an apple` | Low makes "a", Vowel makes "an" |
| 37 | `10 (bin) (hex)` | `2` | 10(bin)->2. 2(hex)->2. (Valid chain) |
| 38 | `hello world (up, 2) (low, 1)` | `HELLO world` | "hello world" -> "HELLO WORLD" -> "HELLO world" |
| 39 | `word . . . word` | `word... word` | Punctuation merging (if treated as individual tokens) |
| 40 | `word ,,,` | `word,,,` | Multiple punctuation marks attached |
| 41 | `boom !!! ?` | `boom!!!?` | Mixed punctuation groups |
| 42 | `' ' ' '` | `'' ''` | Empty quotes |
| 43 | `(hex) (bin) (up)` | `` | Orphan markers (Removed) |
| 44 | `42 (up)` | `42` | Up on number (No effect) |
| 45 | `word (nonsense)` | `word (nonsense)` | Unknown marker (Treat as text) |
| 46 | `(up, -1)` | `` | Negative count (Invalid/Removed) |
| 47 | `a ' big ' apple` | `a 'big' apple` | Quotes separating vowel from target (No "an") |
| 48 | `tok&en (up)` | `TOK&EN` | Special chars in word |
| 49 | `hello (up, 100) world` | `HELLO world` | Count exceeds bounds (Clamp to start) |
| 50 | `sample (cap) (cap) (cap)` | `Sample` | Redundant markers |