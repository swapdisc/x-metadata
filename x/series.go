package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Series() (s Suite) {
	s = Suite{
		Name:         "series",
		Blurb:        `Write a program that will take a string of digits and give you all the possible consecutive number series of length ` + "`" + `n` + "`" + ` in that string.`,
		Description:  seriesDescription(),
		Source:       `A subset of the Problem 8 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=8",
		Expectations: seriesExpectations(),
	}
	return
}

func seriesDescription() string {
	return `For example, the string "01234" has the following 3-digit series:

* 012
* 123
* 234

And the following 4-digit series:

* 0123
* 1234

And if you ask for a 6-digit series from a 5-digit string,
you deserve whatever you get.
`
}

func seriesExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
