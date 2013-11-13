package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func SaddlePoints() (s Suite) {
	s = Suite{
		Name:         "saddle-points",
		Blurb:        `Write a program that detects saddle points in a matrix.`,
		Description:  saddlepointsDescription(),
		Source:       `J Dalbey's Programming Practice problems`,
		SourceUrl:    "http://users.csc.calpoly.edu/~jdalbey/103/Projects/ProgrammingPractice.html",
		Expectations: saddlepointsExpectations(),
	}
	return
}

func saddlepointsDescription() string {
	return `So say you have a matrix like so:

` + "`" + `` + "`" + `` + "`" + `plain
    0  1  2
  |---------
0 | 9  8  7
1 | 5  3  2     <--- saddle point at (1,0)
2 | 6  6  7
` + "`" + `` + "`" + `` + "`" + `

It has a saddle point at (1, 0).

It's called a "saddle point" because it is greater than or equal to every
element in its row and the less than or equal to every element in its column.

A matrix may have zero or more saddle points.

Your code should be able to provide the (possibly empty) list of all the
saddle points for any given matrix.

Note that you may find other definitions of matrix saddle points online, but
the tests for this exercise follow the above unambiguous definition.
`
}

func saddlepointsExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
