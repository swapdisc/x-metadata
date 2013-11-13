package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func ScrabbleScore() (s Suite) {
	s = Suite{
		Name:         "scrabble-score",
		Blurb:        `Write a program that, given a word, computes the scrabble score for that word.`,
		Description:  scrabblescoreDescription(),
		Source:       `Inspired by the Extreme Startup game`,
		SourceUrl:    "https://github.com/rchatley/extreme_startup",
		Expectations: scrabblescoreExpectations(),
	}
	return
}

func scrabblescoreDescription() string {
	return `## Letter Values

You'll need these:

` + "`" + `` + "`" + `` + "`" + `plain
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
` + "`" + `` + "`" + `` + "`" + `

## Examples
"cabbage" should be scored as worth 14 points:

- 3 points for C
- 1 point for A, twice
- 3 points for B, twice
- 2 points for G
- 1 point for E

And to total:

- ` + "`" + `3 + 2*1 + 2*3 + 2 + 1` + "`" + `
- = ` + "`" + `3 + 2 + 6 + 3` + "`" + `
- = ` + "`" + `5 + 9` + "`" + `
- = 14

## Extensions
* You can play a ` + "`" + `:double` + "`" + ` or a ` + "`" + `:triple` + "`" + ` letter.
* You can play a ` + "`" + `:double` + "`" + ` or a ` + "`" + `:triple` + "`" + ` word.
`
}

func scrabblescoreExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
