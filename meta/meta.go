package meta

func All() []suite {
	suites := []suite{
		bob(),
		phrase(),
		anagram(),
		hamming(),
		raindrops(),
	}
	return suites
}
