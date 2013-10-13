package meta

func All() []suite {
	suites := []suite{
		bob(),
		phrase(),
		anagram(),
	}
	return suites
}
