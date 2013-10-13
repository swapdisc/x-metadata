package meta

type phraseOutput map[string]int

func phrase() (s suite) {
	expectations := []expectation{
		expectation{
			"count one word",
			"word",
			phraseOutput{"word": 1},
		},
		expectation{
			"count one of each",
			"one of each",
			phraseOutput{"one": 1, "of": 1, "each": 1},
		},
		expectation{
			"count multiple occurrences",
			"one fish two fish red fish blue fish",
			phraseOutput{"one": 1, "fish": 4, "two": 1, "red": 1, "blue": 1},
		},
		expectation{
			"count everything just once",
			"all the kings horses and all the kings men",
			phraseOutput{"all": 2, "the": 2, "kings": 2, "horses": 1, "and": 1, "men": 1},
		},
		expectation{
			"ignore punctuation",
			"car : carpet as java : javascript!!&@$%^&",
			phraseOutput{"car": 1, "carpet": 1, "as": 1, "java": 1, "javascript": 1},
		},
		expectation{
			"handles cramped lists",
			"one,two,three",
			phraseOutput{"one": 1, "two": 1, "three": 1},
		},
		expectation{
			"include numbers",
			"testing, 1, 2 testing",
			phraseOutput{"testing": 2, "1": 1, "2": 1},
		},
		expectation{
			"normalize case",
			"go Go GO",
			phraseOutput{"go": 3},
		},
		expectation{
			"with apostrophes",
			"First: don't laugh. Then: don't cry.",
			phraseOutput{"first": 1, "don't": 2, "laugh": 1, "then": 1, "cry": 1},
		},
	}

	s = suite{
		Name:         "phrase",
		Expectations: expectations,
	}
	return
}
