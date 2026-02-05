# Golden Tests (Acceptance Criteria)

These tests define the "Done" state. Every test case must pass for the project to be successful.

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

## 2. Additional Golden Tests (Edge Cases)
These are custom tests designed to catch edge cases.

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