package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type suite struct {
	Name         string
	Expectations []expectation
}

func (s suite) filename() (f string) {
	return fmt.Sprintf("exercises/%s.json", s.Name)
}

func (s suite) export() error {
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
	b, err := json.Marshal(s.Expectations)
	return b, err
}
