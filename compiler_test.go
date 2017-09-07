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

func TestCompileMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	originalCompiler := compiler
	defer func() {
		ctrl.Finish()
		compiler = originalCompiler
	}()
	t.Run("CompileNoErrors", func(t *testing.T) {
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
	t.Run("CompileSomeWithErrors", func(t *testing.T) {
		//err = errors.New(fmt.Sprintf("invalid patterns: %s", strings.Join(compileErrors, ",")))
		mockRegexes := []mockRegex{
			{pattern: "error-pattern", rx: nil, err: errors.New("bad pattern!")},
			{pattern: "other-error-pattern", rx: nil, err: errors.New("also bad pattern!")},
			{pattern: "ok-pattern", rx: &regexp.Regexp{}, err: nil},
		}
		expectedErrorMessage := "compilation errors: bad pattern!, also bad pattern!"
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
	rx, _ := CompileAll(patterns)
	fmt.Println(rx[0], reflect.TypeOf(rx[0]))
	fmt.Println(rx[1], reflect.TypeOf(rx[1]))
	// Output:
	// some *regexp.Regexp
	// patterns *regexp.Regexp
}

func ExampleCompileAll2() {
	patterns := []string{
		"(", "ok",
	}
	rx, err := CompileAll(patterns)
	fmt.Println(rx)
	fmt.Println(err)
	// Output:
	// []
	// compilation errors: error parsing regexp: missing closing ): `(`
}