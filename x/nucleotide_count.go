package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func NucleotideCount() (s Suite) {
	s = Suite{
		Name:         "nucleotide-count",
		Blurb:        `Write a class ` + "`" + `DNA` + "`" + ` that takes a DNA string and tells us how many times each nucleotide occurs in the string.`,
		Description:  nucleotidecountDescription(),
		Source:       `The Calculating DNA Nucleotides_problem at Rosalind`,
		SourceUrl:    "http://rosalind.info/problems/dna/",
		Expectations: nucleotidecountExpectations(),
	}
	return
}

func nucleotidecountDescription() string {
	return `DNA is represented by an alphabet of the following symbols: 'A', 'C', 'G', and 'T'.

Each symbol represents a nucleotide, which is a fancy name for the particular molecules that happen to make up a large part of DNA.

Shortest intro to biochemistry EVAR:

* twigs are to birds nests as
* nucleotides are to DNA and RNA as
* amino acids are to proteins as
* sugar is to starch as
* oh crap lipids

I'm not going to talk about lipids because they're crazy complex.

So back to nucleotides.

There are 5 types of nucleotides. 4 of these occur in DNA: ` + "`" + `A` + "`" + `, ` + "`" + `C` + "`" + `, ` + "`" + `G` + "`" + `, and ` + "`" + `T` + "`" + `. 4 occur in RNA: ` + "`" + `A` + "`" + `, ` + "`" + `C` + "`" + `, ` + "`" + `G` + "`" + `, ` + "`" + `U` + "`" + `.

There are no other nucleotides.

`
}

func nucleotidecountExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
