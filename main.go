package main

import (
	"github.com/exercism/xmetadata/meta"
	"github.com/exercism/xmetadata/x"
)

func main() {
	for _, s := range all() {
		s.Export()
	}
}

func all() []meta.Suite {
	suites := []meta.Suite{
		x.Bob(),
		x.Phrase(),
		x.Anagram(),
		x.Hamming(),
		x.Raindrops(),
	}
	return suites
}
