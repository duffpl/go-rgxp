package rgxp

var compiler regexCompiler

func init() {
	compiler = stdRegexCompiler{}
}
