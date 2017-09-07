package rgxp

import (
	"regexp"
	"strings"
	"github.com/pkg/errors"
)

// regexCompiler is wrapper for allowing mock compiler in tests
type regexCompiler interface {
	Compile(pattern string) (*regexp.Regexp, error)
}

// stdRegexCompiler is wrapper for built-in regexp compiler
type stdRegexCompiler struct {}

func (stdRegexCompiler) Compile(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

// CompileAll creates slice of Regexp objects compiled from input slice of patterns.
//
// If during compilation any of patterns caused compilation error result slice will be empty
// and error value will contain list of compiler errors separated by comma
func CompileAll(patterns []string) (result []*regexp.Regexp, err error) {
	cErrors := []string{}
	for _, pattern := range patterns {
		compiled, cErr := compiler.Compile(pattern)
		if cErr != nil {
			cErrors = append(cErrors, cErr.Error())
		}
		if len(cErrors) > 0 {
			continue
		}
		result = append(result, compiled)
	}
	if len(cErrors) > 0 {
		err = errors.New(strings.Join(cErrors, ", "))
	}
	return
}

// MustCompileAll is like CompileAll but panics if any of patterns cannot be parsed.
//
// Can be nicely used along with matchers at cost of sacrificing error checking <sad face>
func MustCompileAll(patterns []string) (result []*regexp.Regexp) {
	result, err := CompileAll(patterns)
	if err != nil {
		panic("rgxp: CompileAll: " + err.Error())
	}
	return
}