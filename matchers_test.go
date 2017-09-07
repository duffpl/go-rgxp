package rgxp

import (
	"regexp"
	"testing"
	"fmt"
)

func TestMatchers(t *testing.T) {
	input := "Hello Everybody!"
	someMatching := []*regexp.Regexp{
		regexp.MustCompile("I'm sorry, Dave\\. I'm afraid I can't do that\\."),
		regexp.MustCompile("Hello"),
	}
	allMatching := []*regexp.Regexp{
		regexp.MustCompile("Hello"),
		regexp.MustCompile("Everybody"),
	}
	noneMatching := []*regexp.Regexp{
		regexp.MustCompile("Zorg"),
	}
	t.Run("MatchAny", func(t *testing.T) {
		if !MatchAny(someMatching, input) {
			t.Fail()
		}
		if !MatchAny(allMatching, input) {
			t.Fail()
		}
		if MatchAny(noneMatching, input) {
			t.Fail()
		}
	})
	t.Run("MatchAll", func(t *testing.T) {
		if MatchAll(someMatching, input) {
			t.Fail()
		}
		if !MatchAll(allMatching, input) {
			t.Fail()
		}
		if MatchAll(noneMatching, input) {
			t.Fail()
		}
	})
}

func ExampleMatchAny() {
	input := "test data"

	fmt.Println(MatchAny([]*regexp.Regexp{
		regexp.MustCompile("only"),
		regexp.MustCompile("one"),
		regexp.MustCompile("will"),
		regexp.MustCompile("match"),
		regexp.MustCompile("test data"),
	}, input))

	fmt.Println(MatchAny([]*regexp.Regexp{
		regexp.MustCompile("no"),
		regexp.MustCompile("match"),
		regexp.MustCompile("here"),
	}, input))

	// Output:
	// true
	// false
}

func ExampleMatchAll() {
	input := "all must match"

	fmt.Println(MatchAll([]*regexp.Regexp{
		regexp.MustCompile("all"),
		regexp.MustCompile("must"),
		regexp.MustCompile("match"),
	}, input))

	fmt.Println(MatchAll([]*regexp.Regexp{
		regexp.MustCompile("all"),
		regexp.MustCompile("must"),
		regexp.MustCompile("match"),
		regexp.MustCompile("no exceptions"),
	}, input))

	// Output:
	// true
	// false
}