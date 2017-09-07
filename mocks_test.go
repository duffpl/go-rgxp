package rgxp

import (
	"regexp"
	"github.com/golang/mock/gomock"
	"github.com/duffpl/go-rgxp/mock"
)

type mockRegex struct {
	pattern string
	rx      *regexp.Regexp
	err     error
}

func getMockCompilerInput(data [] mockRegex) (compilerInput []string) {
	compilerInput = make([]string, len(data))
	for i, testRegex := range data {
		compilerInput[i] = testRegex.pattern
	}
	return
}

func getMockCompilerRegexOutput(data [] mockRegex) (output []*regexp.Regexp) {
	output = make([]*regexp.Regexp, len(data))
	for i, testRegex := range data {
		output[i] = testRegex.rx
	}
	return
}

func getMockCompiler(mockCtrl *gomock.Controller, mockRegexes []mockRegex) *mock.MockregexCompiler {
	mockCompiler := mock.NewMockregexCompiler(mockCtrl)
	for _, mockRegex := range mockRegexes {
		mockCompiler.EXPECT().Compile(mockRegex.pattern).Return(mockRegex.rx, mockRegex.err)
	}
	return mockCompiler
}