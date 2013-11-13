package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func PythagoreanTriplet() (s Suite) {
	s = Suite{
		Name:         "pythagorean-triplet",
		Blurb:        `There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find the product a * b * c.`,
		Description:  pythagoreantripletDescription(),
		Source:       `Problem 9 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=9",
		Expectations: pythagoreantripletExpectations(),
	}
	return
}

func pythagoreantripletDescription() string {
	return `A Pythagorean triplet is a set of three natural numbers, {a, b, c}, for which,

a**2 + b**2 = c**2

For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.
`
}

func pythagoreantripletExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
