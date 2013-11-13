package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func NthPrime() (s Suite) {
	s = Suite{
		Name:         "nth-prime",
		Blurb:        `Write a program that can tell you what the nth prime is.`,
		Description:  nthprimeDescription(),
		Source:       `A variation on Problem 7 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=7",
		Expectations: nthprimeExpectations(),
	}
	return
}

func nthprimeDescription() string {
	return `By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
`
}

func nthprimeExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
