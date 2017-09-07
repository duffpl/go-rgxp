package rgxp

import (
	"regexp"
	"fmt"
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

// CompileAll creates slice of Regexp objects compiled from input slice of patterns
// When during compilation any of patterns caused compilation error result slice will be empty
// and error value will include list of compiler errors
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
		err = errors.New(fmt.Sprintf("compilation errors: %s", strings.Join(cErrors, ", ")))
	}
	return
}
