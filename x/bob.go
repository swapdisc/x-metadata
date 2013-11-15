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
		{
			"stating something",
			"Whatever.",
			"Tom-ay-to, tom-aaaah-to.",
		},
		{
			"shouting",
			"Woah, chill out!",
			"WATCH OUT!",
		},
		{
			"asking a question",
			"Sure.",
			"Does this cryogenic chamber make me look fat?",
		},
		{
			"asking a numeric question",
			"Sure.",
			"You are, what, like 15?",
		},
		{
			"talking forcefully",
			"Whatever.",
			"Let's go make out behind the gym!",
		},
		{
			"using acronyms in regular speech",
			"Whatever.",
			"It's OK if you don't want to go to the DMV.",
		},
		{
			"forceful questions",
			"Woah, chill out!",
			"WHAT THE HELL WERE YOU THINKING?",
		},
		{
			"shouting numbers",
			"Woah, chill out!",
			"1, 2, 3 GO!",
		},
		{
			"only numbers",
			"Whatever.",
			"1, 2, 3",
		},
		{
			"question with only numbers",
			"Sure.",
			"4?",
		},
		{
			"shouting with special characters",
			"Woah, chill out!",
			"ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
		},
		{
			"shouting with no exclamation mark",
			"Woah, chill out!",
			"I HATE YOU",
		},
		{
			"statement containing question mark",
			"Whatever.",
			"Ending with ? means a question.",
		},
		{
			"prattling on",
			"Sure.",
			"Wait! Hang on. Are you going to be OK?",
		},
		{
			"silence",
			"Fine. Be that way!",
			"",
		},
		{
			"prolonged silence",
			"Fine. Be that way!",
			"    ",
		},
		{
			"really prolonged silence",
			"Fine. Be that way!",
			"                 ",
		},
		{
			"multi line trick question",
			"Whatever.",
			"Do I ever change my mind?\nNo.",
		},
	}
	return expectations
}
