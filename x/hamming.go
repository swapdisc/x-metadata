package x

import (
	. "github.com/exercism/xmetadata/meta"
)

type hammingInput struct {
	Strand1 string `json:"strand1"`
	Strand2 string `json:"strand2"`
}

func Hamming() (s Suite) {
	s = Suite{
		Name:         "hamming",
		Blurb:        "Write a program that can calculate the Hamming distance between two DNA strands.",
		Description:  hammingDescription(),
		Source:       "The Calculating Point Mutations problem at Rosalind",
		SourceUrl:    "http://rosalind.info/problems/hamm/",
		Expectations: hammingExpectations(),
	}
	return
}

func hammingDescription() string {
	return `A mutation is simply a mistake that occurs during the creation or copying of a nucleic acid, in particular DNA. Because nucleic acids are vital to cellular functions, mutations tend to cause a ripple effect throughout the cell. Although mutations are technically mistakes, a very rare mutation may equip the cell with a beneficial attribute. In fact, the macro effects of evolution are attributable by the accumulated result of beneficial microscopic mutations over many generations.

The simplest and most common type of nucleic acid mutation is a point mutation, which replaces one base with another at a single nucleotide.

By counting the number of differences between two homologous DNA strands taken from different genomes with a common ancestor, we get a measure of the minimum number of point mutations that could have occurred on the evolutionary path between the two strands.

This is called the 'Hamming distance'

    GAGCCTACTAACGGGAT
    CATCGTAATGACGGCCT
    ^ ^ ^  ^ ^    ^^
`
}
func hammingExpectations() []Expectation {
	expectations := []Expectation{
		Expectation{
			Name:   "no difference between empty strands",
			Input:  hammingInput{"", ""},
			Output: 0,
		},
		Expectation{
			Name:   "no difference between identical strands",
			Input:  hammingInput{"GGACTGA", "GGACTGA"},
			Output: 0,
		},
		Expectation{
			Name:   "complete hamming distance in small strand",
			Input:  hammingInput{"ACT", "GGA"},
			Output: 3,
		},
		Expectation{
			Name:   "hamming distance in off by one strand",
			Input:  hammingInput{"GGACGGATTCTG", "AGGACGGATTCT"},
			Output: 9,
		},
		Expectation{
			Name:   "small hamming distance in middle somewhere",
			Input:  hammingInput{"GGACG", "GGTCG"},
			Output: 1,
		},
		Expectation{
			Name:   "larger distance",
			Input:  hammingInput{"ACCAGGG", "ACTATGG"},
			Output: 2,
		},
		Expectation{
			Name:   "ignores extra length on other strand when longer",
			Input:  hammingInput{"AAACTAGGGG", "AGGCTAGCGGTAGGAC"},
			Output: 3,
		},
		Expectation{
			Name:   "ignores extra length on original strand when longer",
			Input:  hammingInput{"GACTACGGACAGGGTAGGGAAT", "GACATCGCACACC"},
			Output: 5,
		},
	}
	return expectations
}
