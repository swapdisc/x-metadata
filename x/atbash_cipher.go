package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func AtbashCipher() (s Suite) {
	s = Suite{
		Name:         "atbash-cipher",
		Blurb:        `Create an implementation of the atbash cipher, an ancient encryption system created in the Middle East.`,
		Description:  atbashcipherDescription(),
		Source:       `Wikipedia`,
		SourceUrl:    "http://en.wikipedia.org/wiki/Atbash",
		Expectations: atbashcipherExpectations(),
	}
	return
}

func atbashcipherDescription() string {
	return `The Atbash cipher is a simple substitution cipher that relies on transposing all the letters in the alphabet such that the resulting alphabet is backwards. The first letter is replaced with the last letter, the second with the
second-last, and so on.

An Atbash cipher for the Latin alphabet would be as follows:

` + "`" + `` + "`" + `` + "`" + `plain
Plain:  abcdefghijklmnopqrstuvwxyz
Cipher: zyxwvutsrqponmlkjihgfedcba
` + "`" + `` + "`" + `` + "`" + `

It is a very weak cipher because it only has one possible key, and it is a
simple monoalphabetic substitution cipher. However, this may not have been an issue in the cipher's time.

## Examples
- Encoding "test" gives "gvhg"
- Decoding "gvhg" gives "test"
`
}

func atbashcipherExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
