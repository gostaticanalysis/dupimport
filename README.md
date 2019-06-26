# dupimport [![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godoc] [![Travis](https://img.shields.io/travis/gostaticanalysis/dupimport.svg?style=flat-square)][travis] [![Go Report Card](https://goreportcard.com/badge/github.com/gostaticanalysis/dupimport)](https://goreportcard.com/report/github.com/gostaticanalysis/dupimport) [![codecov](https://codecov.io/gh/gostaticanalysis/dupimport/branch/master/graph/badge.svg)](https://codecov.io/gh/gostaticanalysis/dupimport)

`dupimport` finds duplicated imports in same file.

## Install

```sh
$ go get github.com/gostaticanalysis/dupimport/cmd/dupimport
```

## Usage

```sh
$ go vet -vettool=`which dupimport` pkgname
```

<!-- links -->
[godoc]: http://godoc.org/github.com/gostaticanalysis/dupimport
[travis]: https://travis-ci.org/gostaticanalysis/dupimport
