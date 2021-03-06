package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Hexadecimal() (s Suite) {
	s = Suite{
		Name:         "hexadecimal",
		Blurb:        `Write a program that will convert a hexadecimal number, represented as a string (i.e. "10af8c"), to it's decimal equivalent using first principles (i.e. no, you may not use built-in ruby libraries or gems to accomplish the conversion).`,
		Description:  hexadecimalDescription(),
		Source:       `All of Computer Science`,
		SourceUrl:    "http://www.wolframalpha.com/examples/NumberBases.html",
		Expectations: hexadecimalExpectations(),
	}
	return
}

func hexadecimalDescription() string {
	return `On the web we use hexadecimal to represent colors, i.e. green: 008000, teal: 008080, navy: 000080).

The program should consider strings specifying an invalid hexadecimal as the value 0.
`
}

func hexadecimalExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
