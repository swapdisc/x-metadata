package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Allergies() (s Suite) {
	s = Suite{
		Name:         "allergies",
		Blurb:        `Write a program that, given a person's allergy score, can tell them whether or not they're allergic to a given item, and their full list of allergies.`,
		Description:  allergiesDescription(),
		Source:       `Jumpstart Lab Warm-up`,
		SourceUrl:    "http://jumpstartlab.com",
		Expectations: allergiesExpectations(),
	}
	return
}

func allergiesDescription() string {
	return `An allergy test produces a single numeric score which contains the information about all the allergies the person has (that they were tested for).

The list of items (and their value) that were tested are:

* eggs (1)
* peanuts (2)
* shellfish (4)
* strawberries (8)
* tomatoes (16)
* chocolate (32)
* pollen (64)
* cats (128)

So if Tom is allergic to peanuts and chocolate, he gets a score of 34.

Now, given just that score of 34, your program should be able to say:

- Whether Tom is allergic to any one of those allergens listed above.
- All the allergens Tom is allergic to.
`
}

func allergiesExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
