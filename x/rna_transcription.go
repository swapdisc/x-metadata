package x

import (
	. "github.com/exercism/xmetadata/meta"
)

func RnaTranscription() (s Suite) {
	s = Suite{
		Name:         "rna-transcription",
		Blurb:        `Write a program that can translate a given DNA string to the transcribed RNA string corresponding to it.`,
		Description:  rnatranscriptionDescription(),
		Source:       `Rosalind`,
		SourceUrl:    "http://rosalind.info/problems/rna",
		Expectations: rnatranscriptionExpectations(),
	}
	return
}

func rnatranscriptionDescription() string {
	return `Both DNA and RNA strands are a sequence of nucleotides.

The four nucleotides found in DNA are adenine (**A**), cytosine (**C**), guanine (**G**) and thymidine (**T**).

The four nucleotides found in RNA are adenine (**A**), cytosine (**C**), guanine (**G**) and uracil (**U**).

Given a DNA strand, its transcribed RNA strand is formed by replacing all occurrences of thymidine with uracil.

`
}

func rnatranscriptionExpectations() []Expectation {
	expectations := []Expectation{}
	return expectations
}
