package meta

import (
	"testing"
)

var fakeExpectations = map[string]expectation{
	"silence": expectation{
		"really prolonged silence",
		"Fine. Be that way!",
		"                 ",
	},
	"trick": expectation{
		"multi line trick question",
		"Whatever.",
		"Do I ever change my mind?\nNo.",
	},
}

var fakeSuite suite

func init() {
	fakeSuite = suite{
		Name: "stuff",
		Expectations: []expectation{
			fakeExpectations["silence"],
			fakeExpectations["trick"],
		},
	}
}

func TestSuiteFilename(t *testing.T) {
	if fakeSuite.filename() != "exercises/stuff.json" {
		t.Fatalf("Expected filename to be exercises/thing.json, but it was: %s ", fakeSuite.filename())
	}
}

func TestSuiteJSON(t *testing.T) {
	jsonBytes, _ := fakeSuite.JSON()
	json := string(jsonBytes)
	expected := `[{"name":"really prolonged silence","input":"Fine. Be that way!","output":"                 "},{"name":"multi line trick question","input":"Whatever.","output":"Do I ever change my mind?\nNo."}]`
	if json != expected {
		t.Fatalf("Expected JSON to not be %s ", json)
	}
}
