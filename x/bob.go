package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Bob() (s Suite) {
	s = Suite{
		Name:         "bob",
		Blurb:        "Bob is a lackadaisical teenager. In conversation, his responses are very limited.",
		Description:  bobDescription(),
		Source:       "Inspired by the 'Deaf Grandma' exercise in Chris Pine's Learn to Program tutorial.",
		SourceUrl:    "http://pine.fm/LearnToProgram/?Chapter=06",
		Expectations: bobExpectations(),
	}
	return
}

func bobDescription() string {
	return `Bob answers 'Sure.' if you ask him a question.

He answers 'Woah, chill out!' if you yell at him (ALL CAPS).

He says 'Fine. Be that way!' if you address him without actually saying anything.

He answers 'Whatever.' to anything else.
`
}

func bobExpectations() []Expectation {
	expectations := []Expectation{
		Expectation{
			"stating something",
			"Whatever.",
			"Tom-ay-to, tom-aaaah-to.",
		},
		Expectation{
			"shouting",
			"Woah, chill out!",
			"WATCH OUT!",
		},
		Expectation{
			"asking a question",
			"Sure.",
			"Does this cryogenic chamber make me look fat?",
		},
		Expectation{
			"asking a numeric question",
			"Sure.",
			"You are, what, like 15?",
		},
		Expectation{
			"talking forcefully",
			"Whatever.",
			"Let's go make out behind the gym!",
		},
		Expectation{
			"using acronyms in regular speech",
			"Whatever.",
			"It's OK if you don't want to go to the DMV.",
		},
		Expectation{
			"forceful questions",
			"Woah, chill out!",
			"WHAT THE HELL WERE YOU THINKING?",
		},
		Expectation{
			"shouting numbers",
			"Woah, chill out!",
			"1, 2, 3 GO!",
		},
		Expectation{
			"only numbers",
			"Whatever.",
			"1, 2, 3",
		},
		Expectation{
			"question with only numbers",
			"Sure.",
			"4?",
		},
		Expectation{
			"shouting with special characters",
			"Woah, chill out!",
			"ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
		},
		Expectation{
			"shouting with no exclamation mark",
			"Woah, chill out!",
			"I HATE YOU",
		},
		Expectation{
			"statement containing question mark",
			"Whatever.",
			"Ending with ? means a question.",
		},
		Expectation{
			"prattling on",
			"Sure.",
			"Wait! Hang on. Are you going to be OK?",
		},
		Expectation{
			"silence",
			"Fine. Be that way!",
			"",
		},
		Expectation{
			"prolonged silence",
			"Fine. Be that way!",
			"    ",
		},
		Expectation{
			"really prolonged silence",
			"Fine. Be that way!",
			"                 ",
		},
		Expectation{
			"multi line trick question",
			"Whatever.",
			"Do I ever change my mind?\nNo.",
		},
	}
	return expectations
}
