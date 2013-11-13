package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Wordy() (s Suite) {
	s = Suite{
		Name:         "wordy",
		Blurb:        `Write a program that takes a word problem and returns the answer as an integer.`,
		Description:  wordyDescription(),
		Source:       `Inspired by one of the generated questions in the Extreme Startup game.`,
		SourceUrl:    "https://github.com/rchatley/extreme_startup",
		Expectations: wordyExpectations(),
	}
	return
}

func wordyDescription() string {
	return `## Step 1

E.g.

> What is 5 plus 13?

The program should handle large numbers and negative numbers.

Use the tests to drive your solution by deleting the ` + "`" + `skip` + "`" + ` in one test at a time.

## Step 2

E.g.

> What is 5 plus 13?

> What is 7 minus 5?

> What is 6 multiplied by 4?

> What is 25 divided by 5?

## Step 3

E.g.

> What is 5 plus 13 plus 6?

> What is 7 minus 5 minus 1?

> What is 9 minus 3 plus 5?

> What is 3 plus 5 minus 8?

## Step 4

E.g.

> What is 5 plus 13?

> What is 7 minus 5?

> What is 6 times 4?

> What is 25 divided by 5?

> What is 78 plus 5 minus 3?

> What is 18 times 3 plus 16?

> What is 4 times 3 divided by 6?

> What is 4 plus 3 times 2?

## Extensions

Implement questions of the type:

> What is 2 raised to the 5th power?

Remember to write failing tests for this code.
`
}

func wordyExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
