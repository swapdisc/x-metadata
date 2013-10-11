package main

func bob() (s suite) {
	s = suite{Name: "bob"}

	s = s.append(expectation{
		"stating something",
		"Whatever.",
		"Tom-ay-to, tom-aaaah-to.",
	},
		expectation{
      "shouting",
      "Woah, chill out!",
      "WATCH OUT!",
		},
		expectation{
      "asking a question",
      "Sure.",
      "Does this cryogenic chamber make me look fat?",
		},
		expectation{
      "asking a numeric question",
      "Sure.",
      "You are, what, like 15?",
		},
		expectation{
      "talking forcefully",
      "Whatever.",
      "Let's go make out behind the gym!",
		},
		expectation{
      "using acronyms in regular speech",
      "Whatever.",
      "It's OK if you don't want to go to the DMV.",
		},
		expectation{
      "forceful questions",
      "Woah, chill out!",
      "WHAT THE HELL WERE YOU THINKING?",
		},
		expectation{
      "shouting numbers",
      "Woah, chill out!",
      "1, 2, 3 GO!",
		},
		expectation{
      "only numbers",
      "Whatever.",
      "1, 2, 3",
		},
		expectation{
      "question with only numbers",
      "Sure.",
      "4?",
		},
		expectation{
      "shouting with special characters",
      "Woah, chill out!",
      "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
		},
		expectation{
      "shouting with no exclamation mark",
      "Woah, chill out!",
      "I HATE YOU",
		},
		expectation{
      "statement containing question mark",
      "Whatever.",
      "Ending with ? means a question.",
		},
		expectation{
      "prattling on",
      "Sure.",
      "Wait! Hang on. Are you going to be OK?",
		},
		expectation{
      "silence",
      "Fine. Be that way!",
      "",
		},
		expectation{
      "prolonged silence",
      "Fine. Be that way!",
      "    ",
		},
		expectation{
      "really prolonged silence",
      "Fine. Be that way!",
      "                 ",
		},
		expectation{
      "multi line trick question",
      "Whatever.",
      "Do I ever change my mind?\nNo.",
		},
	)
	return s
}
