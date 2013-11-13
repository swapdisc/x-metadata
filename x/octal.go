package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Octal() (s Suite) {
	s = Suite{
		Name:         "octal",
		Blurb:        `Write a program that will convert a octal number, represented as a string (i.e. '1735263'), to its decimal equivalent using first principles (i.e. no, you may not use built-in ruby libraries or gems to accomplish the conversion).`,
		Description:  octalDescription(),
		Source:       `All of Computer Science`,
		SourceUrl:    "http://www.wolframalpha.com/input/?i=base+8",
		Expectations: octalExpectations(),
	}
	return
}

func octalDescription() string {
	return `The program should consider strings specifying an invalid octal as the value 0.

Tests are provided, delete one ` + "`" + `skip` + "`" + ` at a time.
`
}

func octalExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
