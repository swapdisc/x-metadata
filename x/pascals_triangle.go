package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func PascalsTriangle() (s Suite) {
	s = Suite{
		Name:         "pascals-triangle",
		Blurb:        `Write a program that computes Pascal's triangle up to a given number of rows.`,
		Description:  pascalstriangleDescription(),
		Source:       `Pascal's Triangle at Wolfram Math World`,
		SourceUrl:    "http://mathworld.wolfram.com/PascalsTriangle.html",
		Expectations: pascalstriangleExpectations(),
	}
	return
}

func pascalstriangleDescription() string {
	return `In Pascal's Triangle each number is computed by adding the numbers to the
right and left of the current position in the previous row.

` + "`" + `` + "`" + `` + "`" + `plain
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
# ... etc
` + "`" + `` + "`" + `` + "`" + `
`
}

func pascalstriangleExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
