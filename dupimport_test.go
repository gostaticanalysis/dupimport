package dupimport_test

import (
	"testing"

	"github.com/gostaticanalysis/dupimport"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, dupimport.Analyzer, "a", "b", "c", "d", "e", "f", "g")
}

func TestSuggestedFix(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, dupimport.Analyzer, "a", "b", "c", "d", "e", "f", "g")
}
