package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func BeerSong() (s Suite) {
	s = Suite{
		Name:         "beer-song",
		Blurb:        `Write a program which produces the lyrics to that beloved classic, that field-trip favorite: 99 Bottles of Beer on the Wall.`,
		Description:  beersongDescription(),
		Source:       `Learn to Program by Chris Pine`,
		SourceUrl:    "http://pine.fm/LearnToProgram/?Chapter=06",
		Expectations: beersongExpectations(),
	}
	return
}

func beersongDescription() string {
	return `Note that not all verses are identical.

The verse for 1 bottle is as follows:

` + "`" + `` + "`" + `` + "`" + `plain
1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.
` + "`" + `` + "`" + `` + "`" + `

The verse for 0 bottles is as follows:

` + "`" + `` + "`" + `` + "`" + `plain
No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
` + "`" + `` + "`" + `` + "`" + `

You can get multiple verses by passing in the verse to start from, and (optionally) the last verse to include.

## For bonus points

Did you get the tests passing and the code clean? If you want to, these are some additional things you could try:

* Remove as much duplication as you possibly can.
* Optimize for readability, even if it means introducing duplication.
* If you've removed all the duplication, do you have a lot of conditionals? Try applying the _Refactor Conditionals to Strategy Pattern_ refactoring, if it applies in this language. How readable is it?

Then please share your thoughts in a comment on the submission. Did this experiment make the code better? Worse? Did you learn anything from it?

`
}

func beersongExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
