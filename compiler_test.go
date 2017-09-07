package rgxp

import (
	"regexp"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	"testing"
	"errors"
	"fmt"
	"reflect"
)

func TestCompileAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	originalCompiler := compiler
	defer func() {
		ctrl.Finish()
		compiler = originalCompiler
	}()
	t.Run("NoErrors", func(t *testing.T) {
		mockRegexes := []mockRegex{
			{"test-pattern", &regexp.Regexp{}, nil},
			{"other-test-pattern", &regexp.Regexp{}, nil},
		}
		compiler = getMockCompiler(ctrl, mockRegexes)
		compilerInput := getMockCompilerInput(mockRegexes)
		expectedCompilerOutput := getMockCompilerRegexOutput(mockRegexes)
		result, err := CompileAll(compilerInput)
		assert.Equal(t, result, expectedCompilerOutput)
		assert.Nil(t, err)
	})
	t.Run("SomeErrors", func(t *testing.T) {
		//err = errors.New(fmt.Sprintf("invalid patterns: %s", strings.Join(compileErrors, ",")))
		mockRegexes := []mockRegex{
			{pattern: "error-pattern", rx: nil, err: errors.New("bad pattern!")},
			{pattern: "other-error-pattern", rx: nil, err: errors.New("also bad pattern!")},
			{pattern: "ok-pattern", rx: &regexp.Regexp{}, err: nil},
		}
		expectedErrorMessage := "bad pattern!, also bad pattern!"
		compilerInput := getMockCompilerInput(mockRegexes)
		compiler = getMockCompiler(ctrl, mockRegexes)
		result, err := CompileAll(compilerInput)
		assert.Nil(t, result)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func TestMustCompileAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	originalCompiler := compiler
	defer func() {
		compiler = originalCompiler
		ctrl.Finish()
	}()
	t.Run("PanicOnError", func(t *testing.T) {
		mockRegexes := []mockRegex{
			{pattern: "error-pattern", rx: nil, err: errors.New("bad pattern!")},
			{pattern: "ok-pattern", rx: &regexp.Regexp{}, err: nil},
		}
		compilerInput := getMockCompilerInput(mockRegexes)
		compiler = getMockCompiler(ctrl, mockRegexes)
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			MustCompileAll(compilerInput)
		}()
		assert.True(t, panicked)
	})
	t.Run("NoErrors", func(t *testing.T) {
		mockRegexes := []mockRegex{
			{pattern: "ok-pattern", rx: &regexp.Regexp{}, err: nil},
			{pattern: "another-ok-pattern", rx: &regexp.Regexp{}, err: nil},
		}
		compilerInput := getMockCompilerInput(mockRegexes)
		compiler = getMockCompiler(ctrl, mockRegexes)
		result := MustCompileAll(compilerInput)
		assert.Equal(t, result, getMockCompilerRegexOutput(mockRegexes))
	})
}

func ExampleCompileAll() {
	patterns := []string{
		"some", "patterns",
	}

	regexps, _ := CompileAll(patterns)
	fmt.Println(regexps[0], reflect.TypeOf(regexps[0]))
	fmt.Println(regexps[1], reflect.TypeOf(regexps[1]))

	// Output:
	// some *regexp.Regexp
	// patterns *regexp.Regexp
}

// If any of patterns failed compilation result will be empty slice and list of compilation errors separated by comma
func ExampleCompileAll_errors() {
	patterns := []string{
		"(", "ok", ":[",
	}

	regexps, err := CompileAll(patterns)
	fmt.Println(regexps)
	fmt.Println(err)

	// Output:
	// []
	// error parsing regexp: missing closing ): `(`, error parsing regexp: missing closing ]: `[`
}

// Returns slice of compiled Regexps. Can be coupled with matchers nicely
func ExampleMustCompileAll() {
	patterns := []string{
		"all", "should", "match",
	}
	input := "all should match"
	fmt.Println(MatchAll(MustCompileAll(patterns), input))

	patterns = []string{
		"just", "one", "m.*?ch", "is", "enough",
	}
	input = "match me"
	fmt.Println(MatchAny(MustCompileAll(patterns), input))

	// Output:
	// true
	// true
}

// Panics if any of patterns errored during compilation
func ExampleMustCompileAll_error() {

	patterns := []string{
		"ok", "*",
	}
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Something really bad happened")
				fmt.Println(r)
			}
		}()
		MustCompileAll(patterns)
	}()

	// Output:
	// Something really bad happened
	// rgxp: CompileAll: error parsing regexp: missing argument to repetition operator: `*`
}
