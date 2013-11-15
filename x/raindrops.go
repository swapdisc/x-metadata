package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Raindrops() (s Suite) {
	s = Suite{
		Name:         "raindrops",
		Blurb:        "Write a program that converts a number to a string, the contents of which depends on the number's prime factors.",
		Description:  raindropsDescription(),
		Source:       "A variation on a famous interview question intended to weed out the obviously incompetent.",
		SourceUrl:    "http://jumpstartlab.com",
		Expectations: raindropsExpectations(),
	}
	return
}

func raindropsDescription() string {
	return `If the number contains 3 as a prime factor, output 'Pling'.

If the number contains 5 as a prime factor, output 'Plang'.

If the number contains 7 as a prime factor, output 'Plong'.

If the number does not contain 3, 5, or 7 as a prime factor,
just pass the number's digits straight through.

## Examples

28's prime-factorization is 2, 2, 7.
In raindrop-speak, this would be a simple "Plong".

1755 prime-factorization is 3, 3, 3, 5, 13.
In raindrop-speak, this would be a "PlingPlang".

The prime factors of 34 are 2 and 17.
Raindrop-speak doesn't know what to make of that,
so it just goes with the straightforward "34".
`
}

func raindropsExpectations() []Expectation {
	expectations := []Expectation{
		{"1", 1, "1"},
		{"3", 3, "Pling"},
		{"5", 5, "Plang"},
		{"7", 7, "Plong"},
		{"6", 6, "Pling"},
		{"9", 9, "Pling"},
		{"10", 10, "Plang"},
		{"14", 14, "Plong"},
		{"15", 15, "PlingPlang"},
		{"21", 21, "PlingPlong"},
		{"25", 25, "Plang"},
		{"35", 35, "PlangPlong"},
		{"49", 49, "Plong"},
		{"52", 52, "52"},
		{"105", 105, "PlingPlangPlong"},
		{"12121", 12121, "12121"},
	}
	return expectations
}
