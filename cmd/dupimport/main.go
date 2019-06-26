package main

import (
	"github.com/gostaticanalysis/dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimport.Analyzer) }
