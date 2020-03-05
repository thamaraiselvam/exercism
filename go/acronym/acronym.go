// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

type Acronym struct {
	str string
}

func newAcronym(str string) Acronym {
	return Acronym{str: str}
}

// Abbreviate should have a comment documenting it.
func Abbreviate(str string) string {
	acronym := newAcronym(str)
	return acronym.convert()
}

func (acronym Acronym) findWords() [][]string {
	regex, err := regexp.Compile("[a-z|A-Z]+")
	if err != nil {
		os.Exit(1)
	}

	return regex.FindAllStringSubmatch(acronym.str, -1)
}

func (acronym Acronym) isIlEligibleStr(word string) bool {
	return len(word) == 1 && unicode.IsLower(rune(word[0]))
}

func (acronym Acronym) convertLetterToUpper(word string) string {
	firstLetter := string(word[0])
	return strings.ToUpper(firstLetter)
}

func (acronym Acronym) convert() string {
	words := acronym.findWords()

	var result = ""
	for _, word := range words {
		if acronym.isIlEligibleStr(word[0]) {
			continue
		}
		result += acronym.convertLetterToUpper(word[0])
	}

	return result
}
