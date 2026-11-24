# neetcode

My Go solutions to [NeetCode 150](https://neetcode.io/practice/practice/neetcode150) practice problems, each with tests.

## Solutions

### Arrays & Hashing (`arrays/`)

| Problem | Solution | Tests |
| --- | --- | --- |
| Encode and Decode Strings | [encodeDecode.go](arrays/encodeDecode.go) | [encodeDecode_test.go](arrays/encodeDecode_test.go) |
| Longest Consecutive Sequence | [longestConsecutive.go](arrays/longestConsecutive.go) | [longestConsecutive_test.go](arrays/longestConsecutive_test.go) |
| Products of Array Except Self | [productExceptSelf.go](arrays/productExceptSelf.go) | [productExceptSelf_test.go](arrays/productExceptSelf_test.go) |
| Top K Elements in List | [topKFrequency.go](arrays/topKFrequency.go) | [topKFrequency_test.go](arrays/topKFrequency_test.go) |
| Valid Sudoku | [validSudoku.go](arrays/validSudoku.go) | [validSudoku_test.go](arrays/validSudoku_test.go) |

### Two Pointers (`twoPointers/`)

| Problem | Solution | Tests |
| --- | --- | --- |
| Is Palindrome | [isPalindrome.go](twoPointers/isPalindrome.go) | [isPalindrome_test.go](twoPointers/isPalindrome_test.go) |

### Sliding Window (`slidingWindow/`)

| Problem | Solution | Tests |
| --- | --- | --- |
| Buy and Sell Crypto | [maxProfit.go](slidingWindow/maxProfit.go) | [maxProfit_test.go](slidingWindow/maxProfit_test.go) |

Each solution file has a doc comment with the problem statement, constraints and approach.

## Running the tests

Requires Go (see [go.mod](go.mod) for the version).

```bash
go test ./...
```

Tests use [testify](https://github.com/stretchr/testify).
