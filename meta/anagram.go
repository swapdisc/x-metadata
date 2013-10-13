package meta

type anagramInput struct {
	Subject    string   `json:"subject"`
	Candidates []string `json:"candidates"`
}

func anagram() (s suite) {
	expectations := []expectation{
		expectation{
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
		expectation{
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
		expectation{
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
		expectation{
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
		expectation{
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
		expectation{
			"does not confuse different duplicates",
			anagramInput{
				Subject: "galea",
				Candidates: []string{
					"eagle",
				},
			},
			[]string{},
		},
		expectation{
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
		expectation{
			"eliminate anagrams with same checksum",
			anagramInput{
				Subject: "mass",
				Candidates: []string{
					"last",
				},
			},
			[]string{},
		},
		expectation{
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
		expectation{
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
		expectation{
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

	s = suite{
		Name:         "anagram",
		Expectations: expectations,
	}
	return
}
