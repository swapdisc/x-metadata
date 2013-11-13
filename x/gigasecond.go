package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Gigasecond() (s Suite) {
	s = Suite{
		Name:         "gigasecond",
		Blurb:        `Write a program that will calculate the date that someone turned or will celebrate their 1 Gs anniversary.`,
		Description:  gigasecondDescription(),
		Source:       `Chapter 9 in Chris Pine's online Learn to Program tutorial.`,
		SourceUrl:    "http://pine.fm/LearnToProgram/?Chapter=09",
		Expectations: gigasecondExpectations(),
	}
	return
}

func gigasecondDescription() string {
	return `A gigasecond is one billion seconds.
`
}

func gigasecondExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
