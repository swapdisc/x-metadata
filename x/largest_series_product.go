package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func LargestSeriesProduct() (s Suite) {
	s = Suite{
		Name:         "largest-series-product",
		Blurb:        `Write a program that, when given a string of digits, can calculate the largest product for a series of consecutive digits of length n.`,
		Description:  largestseriesproductDescription(),
		Source:       `A variation on Problem 8 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=8",
		Expectations: largestseriesproductExpectations(),
	}
	return
}

func largestseriesproductDescription() string {
	return `For example, for the input ` + "`" + `'0123456789'` + "`" + `, the largest product for a series of 3 digits is 504 (7 * 8 * 9), and the largest product for a series of 5 digits is 15120 (5 * 6 * 7 * 8 * 9).

For the input ` + "`" + `'73167176531330624919225119674426574742355349194934'` + "`" + `, the
largest product is 23520.
`
}

func largestseriesproductExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
