# cyclop

[![Coverage Status](https://coveralls.io/repos/github/bkielbasa/cyclop/badge.svg?branch=master)](https://coveralls.io/github/bkielbasa/cyclop?branch=master)

Cyclop calculates cyclomatic complexities of functions in Go source code.

## Usage

```
go get github.com/bkielbasa/cyclop
```

Options

```
cyclop: calculates cyclomatic complexity

Usage: cyclop [-flag] [package]


Flags:
  -V    print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -json
        emit JSON output
  -maxComplexity int
        max complexity the function can have (default 10)
  -memprofile string
        write memory profile to this file
  -packageAverage float
        max avarage complexity in package
  -skipTests
        should the linter execute on test files as well
  -source
        no effect (deprecated)
  -tags string
        no effect (deprecated)
  -trace string
        write trace log to this file
  -v    no effect (deprecated)

```

Important parameters are:
* `-maxComplexity int` - the max complexity calculated for a single function. `10` by default
* `-packageAvarage float64` - the average cyclomatic complexity for a package. If the value is higher than `0` it will reaise an error if the average will be higher. `0` default. 
* `-skipTests bool` - should checks be executed in tests files. `false` by default
