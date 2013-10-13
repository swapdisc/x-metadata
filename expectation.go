package main

type expectation struct {
	Name   string      `json:"name"`
	Input  string      `json:"input"`
	Output interface{} `json:"output"`
}
