package dupimport_test

import (
	"testing"

	"github.com/gostaticanalysis/dupimport"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, dupimport.Analyzer, "a")
}
