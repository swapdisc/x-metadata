package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Triangle() (s Suite) {
	s = Suite{
		Name:         "triangle",
		Blurb:        `Write a program that can tell you if a triangle is equilateral, isosceles, or scalene.`,
		Description:  triangleDescription(),
		Source:       `The Ruby Koans triangle project, parts 1 & 2`,
		SourceUrl:    "http://rubykoans.com",
		Expectations: triangleExpectations(),
	}
	return
}

func triangleDescription() string {
	return `The program should raise an error if the triangle cannot exist.

Tests are provided, delete one ` + "`" + `skip` + "`" + ` at a time.

## Hint

The sum of the lengths of any two sides of a triangle always exceeds the length of the third side, a principle known as the _triangle inequality_.
`
}

func triangleExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
