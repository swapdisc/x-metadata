package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func WordCount() (s Suite) {
	s = Suite{
		Name:         "word-count",
		Blurb:        `Write a program that given a phrase can count the occurrences of each word in that phrase.`,
		Description:  wordcountDescription(),
		Source:       `The golang tour`,
		SourceUrl:    "http://tour.golang.org",
		Expectations: wordcountExpectations(),
	}
	return
}

func wordcountDescription() string {
	return `For example for the input ` + "`" + `"olly olly in come free"` + "`" + `

` + "`" + `` + "`" + `` + "`" + `plain
olly: 2
in: 1
come: 1
free: 1
` + "`" + `` + "`" + `` + "`" + `

`
}

func wordcountExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
