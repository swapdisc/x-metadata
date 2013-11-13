package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func PalindromeProducts() (s Suite) {
	s = Suite{
		Name:         "palindrome-products",
		Blurb:        `Write a program that can detect palindrome products in a given range.`,
		Description:  palindromeproductsDescription(),
		Source:       `Problem 4 at Project Euler`,
		SourceUrl:    "http://projecteuler.net/problem=4",
		Expectations: palindromeproductsExpectations(),
	}
	return
}

func palindromeproductsDescription() string {
	return `A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.
`
}

func palindromeproductsExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
