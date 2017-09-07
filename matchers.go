package rgxp

import "regexp"

// MatchAny tests input string againts passed slice of Regexp objects
// If any of Regexps passed result will be true
func MatchAny(regexps []*regexp.Regexp, input string) bool {
	for _, rxp := range regexps {
		if rxp.Match([]byte(input)) {
			return true
		}
	}
	return false
}

// MatchAll tests input string againts passed slice of Regexp objects
// Only if all of Regexps passed result will be true
func MatchAll(regexps []*regexp.Regexp, input string) bool {
	for _, rxp := range regexps {
		if !rxp.Match([]byte(input)) {
			return false
		}
	}
	return true
}
