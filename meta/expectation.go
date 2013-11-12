package meta

type Expectation struct {
	Name   string      `json:"name"`
	Input  interface{} `json:"input"`
	Output interface{} `json:"output"`
}
