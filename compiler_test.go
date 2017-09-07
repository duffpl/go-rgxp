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