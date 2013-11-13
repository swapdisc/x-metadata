package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func DifferenceOfSquares() (s Suite) {
	s = Suite{
		Name:         "difference-of-squares",
		Blurb:        `Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.`,
		Description:  differenceofsquaresDescription(),
		Source:       `Problem 6 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=6",
		Expectations: differenceofsquaresExpectations(),
	}
	return
}

func differenceofsquaresDescription() string {
	return `The sum of the squares of the first ten natural numbers is,

1**2 + 2**2 + ... + 10**2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)**2 = 55**2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.
`
}

func differenceofsquaresExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
