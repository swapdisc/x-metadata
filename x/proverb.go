package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Proverb() (s Suite) {
	s = Suite{
		Name:         "proverb",
		Blurb:        `For want of a horseshoe nail, a kingdom was lost, or so the saying goes. Write a program that outputs the full text of this proverbial rhyme.`,
		Description:  proverbDescription(),
		Source:       `Wikipedia`,
		SourceUrl:    "http://en.wikipedia.org/wiki/For_Want_of_a_Nail",
		Expectations: proverbExpectations(),
	}
	return
}

func proverbDescription() string {
	return `For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a horseshoe nail.
`
}

func proverbExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
