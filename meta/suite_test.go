package meta

import (
	"testing"
)

func TestSuiteTitle(t *testing.T) {
	s := Suite{
		Name: "one-two",
	}
	title := s.title()
	if title != "OneTwo" {
		t.Fatalf("Expected title to be OneTwo, but it was: %s ", title)
	}
}

func TestSuiteFilename(t *testing.T) {
	s := Suite{
		Name: "stuff",
	}
	f := s.filename()
	if f != "json/stuff.json" {
		t.Fatalf("Expected filename to be json/stuff.json, but it was: %s ", f)
	}
}

func TestSuiteJSON(t *testing.T) {
	s := Suite{
		Name:        "stuff",
		Blurb:       "Do stuff with programming.",
		Description: "Very straight forward string things.",
		Source:      "The internet.",
		SourceUrl:   "http://example.com",
		Expectations: []Expectation{
			{"one thing", "1", 1},
			{"two things", "2", 2},
		},
	}

	jsonBytes, _ := s.JSON()
	json := string(jsonBytes)
	expected := `{"name":"stuff","blurb":"Do stuff with programming.","description":"Very straight forward string things.","source":"The internet.","source_url":"http://example.com","expectations":[{"name":"one thing","input":"1","output":1},{"name":"two things","input":"2","output":2}]}`
	if json != expected {
		t.Fatalf("Expected JSON to not be %s ", json)
	}
}

func TestSuiteReadme(t *testing.T) {
	s := Suite{
		Name:        "what-if",
		Blurb:       "Do stuff.",
		Description: "Simple.\nVery simple stuff.",
		Source:      "The webz.",
		SourceUrl:   "http://example.org",
	}

	expected := `# WhatIf

Do stuff.

Simple.
Very simple stuff.

## Source

The webz. [view source](http://example.org)`

	if s.Readme() != expected {
		t.Fatalf("Expected README to be %s, not %s ", expected, s.Readme())
	}
}
