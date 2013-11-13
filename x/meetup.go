package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Meetup() (s Suite) {
	s = Suite{
		Name:         "meetup",
		Blurb:        `Calculate the date of meetups.`,
		Description:  meetupDescription(),
		Source:       `Jeremy Hinegardner mentioned a Boulder meetup that happens on the Wednesteenth of every month`,
		SourceUrl:    "https://twitter.com/copiousfreetime",
		Expectations: meetupExpectations(),
	}
	return
}

func meetupDescription() string {
	return `Typically meetups happen on the same day of the week.

Examples are

* the first Monday
* the third Tuesday
* the Wednesteenth
* the last Thursday

There are exactly 7 days that end in '-teenth', therefore one is guaranteed that each day of the week will have a '-teenth' in every month.
`
}

func meetupExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
