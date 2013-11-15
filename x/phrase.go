package x

import (
	. "github.com/exercism/xmetadata/meta"
)

type phraseOutput map[string]int

func Phrase() (s Suite) {
	s = Suite{
		Name:         "phrase",
		Blurb:        "Write a program that given a phrase can count the occurrences of each word in that phrase.",
		Description:  phraseDescription(),
		Source:       "The golang tour",
		SourceUrl:    "http://tour.golang.org",
		Expectations: phraseExpectations(),
	}
	return
}

func phraseDescription() string {
	return `For example for the input 'olly olly in come free'

    olly: 2
    in: 1
    come: 1
    free: 1
`
}

func phraseExpectations() []Expectation {
	expectations := []Expectation{
		{
			"count one word",
			"word",
			phraseOutput{"word": 1},
		},
		{
			"count one of each",
			"one of each",
			phraseOutput{"one": 1, "of": 1, "each": 1},
		},
		{
			"count multiple occurrences",
			"one fish two fish red fish blue fish",
			phraseOutput{"one": 1, "fish": 4, "two": 1, "red": 1, "blue": 1},
		},
		{
			"count everything just once",
			"all the kings horses and all the kings men",
			phraseOutput{"all": 2, "the": 2, "kings": 2, "horses": 1, "and": 1, "men": 1},
		},
		{
			"ignore punctuation",
			"car : carpet as java : javascript!!&@$%^&",
			phraseOutput{"car": 1, "carpet": 1, "as": 1, "java": 1, "javascript": 1},
		},
		{
			"handles cramped lists",
			"one,two,three",
			phraseOutput{"one": 1, "two": 1, "three": 1},
		},
		{
			"include numbers",
			"testing, 1, 2 testing",
			phraseOutput{"testing": 2, "1": 1, "2": 1},
		},
		{
			"normalize case",
			"go Go GO",
			phraseOutput{"go": 3},
		},
		{
			"with apostrophes",
			"First: don't laugh. Then: don't cry.",
			phraseOutput{"first": 1, "don't": 2, "laugh": 1, "then": 1, "cry": 1},
		},
	}
	return expectations
}
