package main

import (
	"github.com/exercism/xmetadata/meta"
)

func main() {
	for _, s := range meta.All() {
		s.Export()
	}
}
