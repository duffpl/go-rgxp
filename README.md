# rgxp

rgpx adds some utility methods extending Go's built-in regexp pacakge.

Currently there are just 4 methods

- **CompileAll** accepts slice of patterns, compiles them and returns slice of Regexp objects  or list of compiler errors if something went wrong
- **MustCompileAll** behaves like CompileAll but returns only slice of Regexps or panics if compilation failed
- **MatchAny** test input string against slice of Regexps and if any of them pass result is true
- **MatchAll** does same thing but true is returned only if all Regexps have passed

[Check out docs for more detailed description and examples](https://godoc.org/github.com/duffpl/go-rgxp)
