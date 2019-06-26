package dupimport

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  Doc,
	Run:  run,
}

const Doc = "dupimport is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	paths := map[string]bool{}
	for _, f := range pass.Files {
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}
			if paths[path] {
				pass.Reportf(ip.Pos(), "%s is duplicated import", path)
			} else {
				paths[path] = true
			}
		}
	}
	return nil, nil
}
