package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func OcrNumbers() (s Suite) {
	s = Suite{
		Name:         "ocr-numbers",
		Blurb:        `Write a program that, given a 3 x 4 grid of pipes, underscores, and spaces, can determine which number is represented, or whether it is garbled.`,
		Description:  ocrnumbersDescription(),
		Source:       `Inspired by the Bank OCR kata`,
		SourceUrl:    "http://codingdojo.org/cgi-bin/wiki.pl?KataBankOCR",
		Expectations: ocrnumbersExpectations(),
	}
	return
}

func ocrnumbersDescription() string {
	return `## Step 1

A simple binary font has been constructed using only pipes and underscores.

The number is four rows high, three columns wide:

    _   #
   | |  # zero.
   |_|  #
        # the fourth row is always blank

        #
     |  # one.
     |  #
        # (blank fourth row)

Write a program that, given a 3 x 4 grid of pipes, underscores, and spaces,
can determine whether the the grid represents a zero, a one, or garble.

Anything else is considered garble, and can be represented with a '?'

## Step 2

A simple numeric font has been constructed using only pipes and underscores.

The number consists of four rows high, three columns wide:

    _  _     _  _  _  _  _  _  #
  | _| _||_||_ |_   ||_||_|| | # decimal numbers.
  ||_  _|  | _||_|  ||_| _||_| #
                               # fourth line is always blank

There may be several numbers in the input text, one per line.
`
}

func ocrnumbersExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
