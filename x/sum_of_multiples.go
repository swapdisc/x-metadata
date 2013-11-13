package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func SumOfMultiples() (s Suite) {
	s = Suite{
		Name:         "sum-of-multiples",
		Blurb:        `Write a program that, given a number, can find the sum of all the multiples of 3 or 5 up to and including that number.`,
		Description:  sumofmultiplesDescription(),
		Source:       `A variation on Problem 1 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=1",
		Expectations: sumofmultiplesExpectations(),
	}
	return
}

func sumofmultiplesDescription() string {
	return `If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Allow the program to be configured to find the sum of multiples of numbers other than 3 and 5.
`
}

func sumofmultiplesExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
