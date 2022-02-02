package pattern

import "github.com/prasnitt/stack-via-golang/stack/chars"

const (
	invalidChar = -1
)

// startEndMap will tell the ending char
var startEndMap map[rune]rune = map[rune]rune{
	[]rune("{")[0]: []rune("}")[0],
	[]rune("[")[0]: []rune("]")[0],
	[]rune("(")[0]: []rune(")")[0],
}

// isStartChar is a private function that checks if a character is starting character
func isStartChar(look rune) bool {
	if _, ok := startEndMap[look]; ok {
		return true
	}
	return false
}

// isEndChar is a private function that checks if a character is ending character
func isEndChar(look rune) bool {
	for _, end := range startEndMap {
		if end == look {
			return true
		}
	}
	return false
}

// endingChar is a private function that returns corresponding ending character
func endingChar(start rune) rune {
	if end, ok := startEndMap[start]; ok {
		return end
	}
	return invalidChar
}

// IsValid will check if given pattern is valid or not
func IsValid(pat string) bool {
	var charStack chars.Stack

	for _, c := range pat {
		if isStartChar(c) {
			// for start character push corresponding ending character to stack
			charStack.Push(endingChar(c))
		} else if isEndChar(c) {
			// for end character comes, then get the expecting end character from stack
			// and match it with this character
			expected, err := charStack.Pop()

			// If stack was empty it means, no corresponding staring character found
			if err != nil {
				return false
			}

			// if expected value does not match, means this is wrong end character
			if expected != c {
				return false
			}
		} else {
			// TODO: ignoring other characters other than start and end character
		}
	}

	// If stack is still not empty it means not all ending characters present in pattern string
	return charStack.IsEmpty()
}
