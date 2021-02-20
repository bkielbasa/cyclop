# cyclop

[![Coverage Status](https://coveralls.io/repos/github/bkielbasa/cyclop/badge.svg?branch=master)](https://coveralls.io/github/bkielbasa/cyclop?branch=master)

Cyclop calculates cyclomatic complexities of functions or packages in Go source code.

## Why cyclop?

Cyclop, compared to [other alternative](https://github.com/fzipp/gocyclo), calculates both function and package cyclomatic complexity.

## Usage

```
go get github.com/bkielbasa/cyclop/cmd/cyclop

cyclop .
```

Available parameters:
* `-maxComplexity int` - the max complexity calculated for a single function. `10` by default
* `-packageAverage float64` - the average cyclomatic complexity for a package. If the value is higher than `0` it will reaise an error if the average will be higher. `0` default. 
* `-skipTests bool` - should checks be executed in tests files. `false` by default
