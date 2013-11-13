package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func ParallelLetterFrequency() (s Suite) {
	s = Suite{
		Name:         "parallel-letter-frequency",
		Blurb:        `Write a program that, counts the frequency of letters in texts using parallel computation.`,
		Description:  parallelletterfrequencyDescription(),
		Source:       ``,
		SourceUrl:    "",
		Expectations: parallelletterfrequencyExpectations(),
	}
	return
}

func parallelletterfrequencyDescription() string {
	return `Parallelism is about doing things in parallel that can also be done
sequentially. A common example is counting the frequency of letters. Create a
function that returns the total frequency of each letter in a list of texts and
that employs parallelism.
`
}

func parallelletterfrequencyExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
