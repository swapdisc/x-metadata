package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func SecretHandshake() (s Suite) {
	s = Suite{
		Name:         "secret-handshake",
		Blurb:        `Write a program that will take a decimal number, and convert it to the appropriate sequence of events for a secret handshake.`,
		Description:  secrethandshakeDescription(),
		Source:       `Bert, in Mary Poppins`,
		SourceUrl:    "http://www.imdb.com/character/ch0011238/quotes",
		Expectations: secrethandshakeExpectations(),
	}
	return
}

func secrethandshakeDescription() string {
	return `> There are 10 types of people in the world: Those who understand binary, and those who don't.

You and your fellow cohort of those in the "know" when it comes to binary decide to come up with a secret "handshake".

` + "`" + `` + "`" + `` + "`" + `
1 = wink
10 = double blink
100 = close your eyes
1000 = jump


10000 = Reverse the order of the operations in the secret handshake.
` + "`" + `` + "`" + `` + "`" + `

` + "`" + `` + "`" + `` + "`" + `
handshake = SecretHandshake.new 9
handshake.commands # => ["wink","jump"]

handshake = SecretHandshake.new "11001"
handshake.commands # => ["jump","wink"]
` + "`" + `` + "`" + `` + "`" + `

The program should consider strings specifying an invalid binary as the value 0.
`
}

func secrethandshakeExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
