package dupimport

import (
	"fmt"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

// Analyzer finds duplicated imports in same file.
var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  doc,
	Run:  run,
}

const doc = "dupimport finds duplicated imports in same file"

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		paths := map[string]bool{}
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}
			if paths[path] {
				// Calculate the range covering the entire import spec line.
				// This removes the full line including leading whitespace and trailing newline.
				file := pass.Fset.File(ip.Pos())
				line := file.Line(ip.Pos())
				lineStart := file.LineStart(line)
				var lineEnd token.Pos
				if line < file.LineCount() {
					lineEnd = file.LineStart(line + 1)
				} else {
					lineEnd = token.Pos(file.Base() + file.Size())
				}

				pass.Report(analysis.Diagnostic{
					Pos:     ip.Pos(),
					End:     ip.End(),
					Message: fmt.Sprintf("%s is duplicated import", path),
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: fmt.Sprintf("Remove duplicated import of %s", path),
							TextEdits: []analysis.TextEdit{
								{Pos: lineStart, End: lineEnd, NewText: []byte{}},
							},
						},
					},
				})
			} else {
				paths[path] = true
			}
		}
	}
	return nil, nil
}
