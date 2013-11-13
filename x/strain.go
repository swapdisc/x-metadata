package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Strain() (s Suite) {
	s = Suite{
		Name:         "strain",
		Blurb:        `Implement the ` + "`" + `keep` + "`" + ` and ` + "`" + `discard` + "`" + ` operation on collections. Given a collection and a predicate on the collection's elements, ` + "`" + `keep` + "`" + ` returns a new collection containing those elements where the predicate is true, while ` + "`" + `discard` + "`" + ` returns a new collection containing those elements where the predicate is false.`,
		Description:  strainDescription(),
		Source:       `Conversation with James Edward Gray II`,
		SourceUrl:    "https://twitter.com/jeg2",
		Expectations: strainExpectations(),
	}
	return
}

func strainDescription() string {
	return `For example, given the collection of numbers:

- 1, 2, 3, 4, 5

And the predicate:

- is the number even?

Then your ` + "`" + `keep` + "`" + ` operation should produce:

- 2, 4

While your ` + "`" + `discard` + "`" + ` operation should produce:

- 1, 3, 5

Note that the union of ` + "`" + `keep` + "`" + ` and ` + "`" + `discard` + "`" + ` is all the elements.

## Restrictions

Keep your hands off that filter/reject/whatchamacallit functionality
provided by your standard library!
Solve this one yourself using other basic tools instead.
`
}

func strainExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
