package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func PigLatin() (s Suite) {
	s = Suite{
		Name:         "pig-latin",
		Blurb:        `Implement a program that translates from English to Pig Latin`,
		Description:  piglatinDescription(),
		Source:       `The Pig Latin exercise at Test First Teaching by Ultrasaurus`,
		SourceUrl:    "https://github.com/ultrasaurus/test-first-teaching/blob/master/learn_ruby/pig_latin/",
		Expectations: piglatinExpectations(),
	}
	return
}

func piglatinDescription() string {
	return `Pig Latin is a made-up children's language that's intended to be confusing. It obeys a few simple rules (below) but when it's spoken quickly it's really difficult for non-children (and non-native speakers) to understand.

* **Rule 1**: If a word begins with a vowel sound, add an "ay" sound to the end of the word.
* **Rule 2**: If a word begins with a consonant sound, move it to the end of the word, and then add an "ay" sound to the end of the word.

There are a few more rules for edge cases, and there are regional variants too.

See <http://en.wikipedia.org/wiki/Pig_latin> for more details.
`
}

func piglatinExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
