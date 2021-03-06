// Package string2 is the supplement of the standard library of strings.
package string2

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

var (
	leftDelimiter  = "{"
	rightDelimiter = "}"
)

// SetFmtDelimiter sets the delimiters which are used by KvFmt.
//
// The left delimiter is "{", and the right delimiter is "}".
func SetFmtDelimiter(left, right string) {
	if left == "" || right == "" {
		panic("The arguments cannot be empty")
	}
	leftDelimiter = left
	rightDelimiter = right
}

// KvFmt formats the string like the key-value method format of str in Python,
// which the placeholder is appointed by the key name of the values.
//
// Notice: the formatter will use %v to convert the value of the key to string.
// The delimiters are "{" and "}" by default, and you can reset them by the
// function SetFmtDelimiter.
func KvFmt(s string, values map[string]interface{}) string {
	for key, value := range values {
		key = fmt.Sprintf("%s%s%s", leftDelimiter, key, rightDelimiter)
		s = strings.Replace(s, key, fmt.Sprintf("%v", value), -1)
	}
	return s
}

// SplitSpace splits the string of s by the whitespace, which is equal to
// str.split() in Python.
//
// Notice: SplitSpace(s) == Split(s, unicode.IsSpace).
func SplitSpace(s string) []string {
	return SplitSpaceN(s, -1)
}

// SplitString splits the string of s by sep, but is not the same as
// strings.Split(), which the rune in sep arbitrary combination. For example,
// SplitString("abcdefg-12345", "3-edc") == []string{"ab", "fg", "12", "45"}.
func SplitString(s string, sep string) []string {
	return SplitStringN(s, sep, -1)
}

// Split splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator.
func Split(s string, filter func(c rune) bool) []string {
	return SplitN(s, filter, -1)
}

// SplitSpaceN is the same as SplitStringN, but the whitespace.
func SplitSpaceN(s string, number int) []string {
	return SplitN(s, unicode.IsSpace, number)
}

// SplitStringN is the same as SplitN, but the separator is the string of sep.
func SplitStringN(s string, sep string, number int) []string {
	return SplitN(s, func(c rune) bool {
		for _, r := range sep {
			if r == c {
				return true
			}
		}
		return false
	}, number)
}

// SplitN splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator, but only number times.
//
// If number is equal to 0, don't split; greater than 0, only split number times;
// less than 0, don't limit. If the leading rune is the separator, it doesn't
// consume the split number.
//
// Notice: The result does not have the element of nil.
func SplitN(s string, filter func(c rune) bool, number int) []string {
	if number == 0 {
		return []string{s}
	}

	j := 0
	for i, c := range s {
		if filter(c) {
			j = i
		} else {
			break
		}
	}
	if j != 0 {
		s = s[j+1:]
	}

	if len(s) == 0 {
		return nil
	}

	results := make([]string, 0)
	buf := bytes.NewBuffer(nil)
	isNew := false
	for i, c := range s {
		if filter(c) {
			isNew = true
			continue
		}

		if isNew {
			results = append(results, buf.String())
			buf = bytes.NewBuffer(nil)
			isNew = false
			number--
			if number == 0 {
				buf.WriteString(s[i:])
				break
			}
		}

		buf.WriteRune(c)
	}

	last := buf.String()
	if len(last) > 0 {
		results = append(results, last)
	}

	return results
}
