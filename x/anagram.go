package x

import (
	. "github.com/exercism/xmetadata/meta"
)

type anagramInput struct {
	Subject    string   `json:"subject"`
	Candidates []string `json:"candidates"`
}

func Anagram() (s Suite) {
	s = Suite{
		Name:         "anagram",
		Blurb:        "Write a program that, given a word and a list of possible anagrams, selects the correct sublist.",
		Description:  anagramDescription(),
		Source:       "Inspired by the Extreme Startup game",
		SourceUrl:    "https://github.com/rchatley/extreme_startup",
		Expectations: anagramExpectations(),
	}
	return
}

func anagramDescription() string {
	return "Given the subject 'listen' and a list of candidates such as 'enlists' 'google' 'inlets' 'banana' the program should return a list containing 'inlets'."
}

func anagramExpectations() []Expectation {
	expectations := []Expectation{
		{
			"no matches",
			anagramInput{
				Subject: "diaper",
				Candidates: []string{
					"hello",
					"world",
					"zombies",
					"pants",
				},
			},
			[]string{},
		},
		{
			"detect simple anagram",
			anagramInput{
				Subject: "ant",
				Candidates: []string{
					"tan",
					"stand",
					"at",
				},
			},
			[]string{"tan"},
		},
		{
			"detect another simple anagram",
			anagramInput{
				Subject: "listen",
				Candidates: []string{
					"enlists",
					"google",
					"inlets",
					"banana",
				},
			},
			[]string{"inlets"},
		},
		{
			"detect multiple anagrams",
			anagramInput{
				Subject: "master",
				Candidates: []string{
					"stream",
					"pigeon",
					"maters",
				},
			},
			[]string{"maters", "stream"},
		},
		{
			"detect multiple anagrams again",
			anagramInput{
				Subject: "allergy",
				Candidates: []string{
					"gallery",
					"ballerina",
					"regally",
					"clergy",
					"largely",
					"leading",
				},
			},
			[]string{"gallery", "largely", "regally"},
		},
		{
			"does not confuse different duplicates",
			anagramInput{
				Subject: "galea",
				Candidates: []string{
					"eagle",
				},
			},
			[]string{},
		},
		{
			"identical word is not anagram",
			anagramInput{
				Subject: "corn",
				Candidates: []string{
					"corn",
					"dark",
					"Corn",
					"rank",
					"CORN",
					"cron",
					"park",
				},
			},
			[]string{},
		},
		{
			"eliminate anagrams with same checksum",
			anagramInput{
				Subject: "mass",
				Candidates: []string{
					"last",
				},
			},
			[]string{},
		},
		{
			"eliminate anagram subsets",
			anagramInput{
				Subject: "good",
				Candidates: []string{
					"dog",
					"goody",
				},
			},
			[]string{},
		},
		{
			"subjects are case insensitive",
			anagramInput{
				Subject: "Orchestra",
				Candidates: []string{
					"cashregiser",
					"carthorse",
					"radishes",
				},
			},
			[]string{"carthorse"},
		},
		{
			"candidates are case insensitive",
			anagramInput{
				Subject: "orchestra",
				Candidates: []string{
					"cashregiser",
					"Carthorse",
					"radishes",
				},
			},
			[]string{"carthorse"},
		},
	}
	return expectations
}
