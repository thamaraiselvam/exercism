// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

var replies = map[string]string{
	"DEFAULT":  "Whatever.",
	"NOTHING": "Fine. Be that way!",
	"QUESTION": "Sure.",
	"SHOUTING": "Calm down, I know what I'm doing!",
	"YELLING_QUESTION":  "Whoa, chill out!",
}

type Remark struct {
	remark string
}

func newRemark(s string) Remark {
	return Remark{remark: strings.TrimSpace(s)}
}

func Hey(msg string) string {

	remark := newRemark(msg)

	switch {
	case remark.isShouting():
		return replies["SHOUTING"]
	case remark.isYellingQuestion():
		return replies["YELLING_QUESTION"]
	case remark.isQuestion():
		return replies["QUESTION"]
	case remark.saidNothing():
		return replies["NOTHING"]
	default:
		return replies["DEFAULT"]
	}

}

func (r Remark) isShouting() bool {
	return r.isYellingQuestion() && r.isQuestion()
}

func (r Remark) saidNothing() bool {
	return len(r.remark) == 0
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isYellingQuestion() bool {

	hasLetter := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUpCased := strings.ToUpper(r.remark) == r.remark

	return hasLetter && isUpCased
}
