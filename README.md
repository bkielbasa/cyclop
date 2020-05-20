# cyclop

[![Coverage Status](https://coveralls.io/repos/github/bkielbasa/cyclop/badge.svg?branch=master)](https://coveralls.io/github/bkielbasa/cyclop?branch=master)

Cyclop calculates cyclomatic complexities of functions in Go source code.

## Usage

```
go get github.com/bkielbasa/cyclop
```

Options

```
  -fail-from int
        returns the program with non-zero result if find at least one function with complexity higher than N (default 10)
  -no-tests
        should ignore test files (default true)
  -short-avg
        displays only average complexity in short format
  -top int
        displays top N functions with the biggest complexity
  -total
        displays total coverage

```

Examples

```
cyclop .
cyclop -no-tests=true .
cyclop -top=10 .
```
