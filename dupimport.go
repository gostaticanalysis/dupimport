package dupimport

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path"
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

// importName returns the local name used to refer to the import spec.
// If the import has an explicit alias, it returns that; otherwise it returns
// the base of the import path (e.g. "fmt" for "fmt", "http" for "net/http").
func importName(spec *ast.ImportSpec, importPath string) string {
	if spec.Name != nil {
		return spec.Name.Name
	}
	return path.Base(importPath)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		specs := map[string]*ast.ImportSpec{}
		for _, ip := range f.Imports {
			importPath, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}

			first, isDup := specs[importPath]
			if !isDup {
				specs[importPath] = ip
				continue
			}

			file := pass.Fset.File(ip.Pos())
			line := file.Line(ip.Pos())
			lineStart := file.LineStart(line)
			var lineEnd token.Pos
			if line < file.LineCount() {
				lineEnd = file.LineStart(line + 1)
			} else {
				lineEnd = token.Pos(file.Base() + file.Size())
			}

			edits := []analysis.TextEdit{
				{Pos: lineStart, End: lineEnd, NewText: []byte{}},
			}

			// Find the Object for the duplicate import.
			var dupObj types.Object
			if ip.Name != nil {
				dupObj = pass.TypesInfo.ObjectOf(ip.Name)
			} else {
				dupObj = pass.TypesInfo.Implicits[ip]
			}

			replacement := importName(first, importPath)

			// Find all usages of the duplicate import and create replacement edits.
			if dupObj != nil {
				for ident, obj := range pass.TypesInfo.Uses {
					if obj == dupObj {
						edits = append(edits, analysis.TextEdit{
							Pos:     ident.Pos(),
							End:     ident.End(),
							NewText: []byte(replacement),
						})
					}
				}
			}

			pass.Report(analysis.Diagnostic{
				Pos:     ip.Pos(),
				End:     ip.End(),
				Message: fmt.Sprintf("%s is duplicated import", importPath),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message:   fmt.Sprintf("Remove duplicated import of %s", importPath),
						TextEdits: edits,
					},
				},
			})
		}
	}
	return nil, nil
}
