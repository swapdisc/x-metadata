package meta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type suite struct {
	Name         string        `json:"name"`
	Blurb        string        `json:"blurb"`
	Description  string        `json:"description"`
	Source       string        `json:"source"`
	SourceUrl    string        `json:"source_url"`
	Expectations []expectation `json:"expectations"`
}

func (s suite) filename() (f string) {
	return fmt.Sprintf("json/%s.json", s.Name)
}

func (s suite) Export() error {
	bytes, err := s.JSON()
	if err != nil {
		fmt.Errorf("Could not make json out of expectations in %s, %v", s.Name, err)
		return err
	}

	err = ioutil.WriteFile(s.filename(), bytes, 0644)
	if err != nil {
		fmt.Errorf("Could not write file, %v", err)
	}
	return err
}

func (s suite) JSON() ([]byte, error) {
	b, err := json.Marshal(s)
	return b, err
}
