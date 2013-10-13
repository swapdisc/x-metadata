package meta

type expectation struct {
	Name   string      `json:"name"`
	Input  interface{} `json:"input"`
	Output interface{} `json:"output"`
}
