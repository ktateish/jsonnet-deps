package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/go-jsonnet"
)

//
// Functions for jsonnet-deps
//
type dependLoggingImporter struct {
	deps []string
	next jsonnet.Importer
}

func newDependLoggingImporter(next jsonnet.Importer) *dependLoggingImporter {
	return &dependLoggingImporter{
		deps: []string{},
		next: next,
	}
}

func (dli *dependLoggingImporter) Import(from, path string) (jsonnet.Contents, string, error) {
	dir := filepath.Dir(from)
	if dir[len(dir)-1] != '/' {
		dir += "/"
	}
	if strings.HasPrefix(dir, "./") {
		dir = dir[2:]
	}
	dli.deps = append(dli.deps, fmt.Sprintf("%s%s", dir, path))
	return dli.next.Import(from, path)
}

func (dli *dependLoggingImporter) dependencyList() []string {
	return dli.deps
}
