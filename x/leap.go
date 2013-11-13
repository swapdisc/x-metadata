package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func Leap() (s Suite) {
	s = Suite{
		Name:         "leap",
		Blurb:        `Write a program that will take a year and report if it is a leap year.`,
		Description:  leapDescription(),
		Source:       `JavaRanch Cattle Drive, exercise 3`,
		SourceUrl:    "http://www.javaranch.com/leap.jsp",
		Expectations: leapExpectations(),
	}
	return
}

func leapDescription() string {
	return `The tricky thing here is that a leap year occurs:

` + "`" + `` + "`" + `` + "`" + `plain
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    except every year that is evenly divisible by 400.
` + "`" + `` + "`" + `` + "`" + `

For example, 1997 is not a leap year, but 1996 is.
1900 is not a leap year, but 2000 is.

## Notes

For a delightful, four minute explanation of the whole leap year phenomenon, go watch [this youtube video](http://www.youtube.com/watch?v=xX96xng7sAE).
`
}

func leapExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
