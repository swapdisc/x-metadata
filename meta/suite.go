package meta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Suite struct {
	Name         string        `json:"name"`
	Blurb        string        `json:"blurb"`
	Description  string        `json:"description"`
	Source       string        `json:"source"`
	SourceUrl    string        `json:"source_url"`
	Expectations []Expectation `json:"expectations"`
}

func (s Suite) filename() (f string) {
	return fmt.Sprintf("json/%s.json", s.Name)
}

func (s Suite) title() string {
	t := strings.Split(s.Name, string("-"))
	res := make([]string, len(t))
	for _, part := range t {
		res = append(res, strings.Title(part))
	}
	return strings.Join(res, "")
}

func (s Suite) Readme() (md string) {
	md = "# %s\n\n%s\n\n%s\n\n## Source\n\n%s [view source](%s)"
	return fmt.Sprintf(md, s.title(), s.Blurb, s.Description, s.Source, s.SourceUrl)
}

func (s Suite) Export() error {
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

func (s Suite) WriteReadme() error {
	file := fmt.Sprintf("readme/%s.md", s.title())
	err := ioutil.WriteFile(file, []byte(s.Readme()), 0644)
	if err != nil {
		fmt.Errorf("Could not write file, %v", err)
	}
	return err
}

func (s Suite) JSON() ([]byte, error) {
	b, err := json.Marshal(s)
	return b, err
}
