package main

import (
	"github.com/exercism/xmetadata/meta"
	"github.com/exercism/xmetadata/x"
)

func main() {
	for _, s := range all() {
		s.Export()
		s.WriteReadme()
	}
}

func all() []meta.Suite {
	suites := []meta.Suite{
		x.Bob(),
		x.Phrase(),
		x.Anagram(),
		x.Hamming(),
		x.Raindrops(),
		x.Accumulate(),
		x.Allergies(),
		x.AtbashCipher(),
		x.BankAccount(),
		x.BeerSong(),
		x.Binary(),
		x.BinarySearchTree(),
		x.CryptoSquare(),
		x.DifferenceOfSquares(),
		x.Etl(),
		x.Gigasecond(),
		x.GoCounting(),
		x.GradeSchool(),
		x.Grains(),
		x.Hexadecimal(),
		x.KindergartenGarden(),
		x.LargestSeriesProduct(),
		x.Leap(),
		x.LinkedList(),
		x.Luhn(),
		x.Matrix(),
		x.Meetup(),
		x.NthPrime(),
		x.NucleotideCount(),
		x.OcrNumbers(),
		x.Octal(),
		x.PalindromeProducts(),
		x.ParallelLetterFrequency(),
		x.PascalsTriangle(),
		x.PhoneNumber(),
		x.PigLatin(),
		x.PointMutations(),
		x.PrimeFactors(),
		x.PythagoreanTriplet(),
		x.QueenAttack(),
		x.RnaTranscription(),
		x.RobotName(),
		x.RobotSimulator(),
		x.RomanNumerals(),
		x.SaddlePoints(),
		x.Say(),
		x.ScrabbleScore(),
		x.SecretHandshake(),
		x.Series(),
		x.SgfParsing(),
		x.Sieve(),
		x.SimpleCipher(),
		x.SimpleLinkedList(),
		x.SpaceAge(),
		x.Strain(),
		x.SumOfMultiples(),
		x.Triangle(),
		x.Trinary(),
		x.WordCount(),
		x.Wordy(),
		x.Zipper(),
	}
	return suites
}
