// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.Trim(remark, " ")
	if strings.Trim(remark, " \t\n\r") == "" {
		return "Fine. Be that way!"
	}
	if strings.ToUpper(remark) == remark && remark[len(remark)-1:] == "?" && IsLetter(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if remark[len(remark)-1:] == "?" {
		return "Sure."
	}
	if strings.ToUpper(remark) == remark && IsLetter(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
