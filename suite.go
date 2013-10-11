package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type suite struct {
	Name string
	Expectations []expectation
}

func (s suite) append(e ...expectation) (cp suite) {
	s.Expectations = append(s.Expectations, e...)
	cp = s
	return
}

func (s suite) filename() (f string) {
	f = fmt.Sprintf("%s/%s.json", "exercises", s.Name)
	return
}

func (s suite) export() {
	bytes, err := json.Marshal(s.Expectations)
	if err != nil {
		fmt.Errorf("Could not make json out of expectations in %s, %v", s.Name, err)
	}

	err = ioutil.WriteFile(s.filename(), bytes, 0644)
	if err != nil {
		fmt.Errorf("Could not write file, %v", err)
	}
}

