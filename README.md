# go-jdn

go-jdn is the tool of Julian Day Number.

[![GoDoc](https://godoc.org/github.com/DestinyLab/go-jdn?status.svg)](https://godoc.org/github.com/DestinyLab/go-jdn) [![Go Report Card](https://goreportcard.com/badge/github.com/DestinyLab/go-jdn)](https://goreportcard.com/report/github.com/DestinyLab/go-jdn) [![Build Status](https://travis-ci.org/DestinyLab/go-jdn.svg?branch=master)](https://travis-ci.org/DestinyLab/go-jdn) [![Coverage Status](https://coveralls.io/repos/github/DestinyLab/go-jdn/badge.svg?branch=master)](https://coveralls.io/github/DestinyLab/go-jdn?branch=master)

## Installation

```
go get -u github.com/DestinyLab/go-jdn
```

## Usage

```go
package main

import (
  "fmt"
  "time"

  "github.com/DestinyLab/go-jdn"
)

func main() {
  t1 := time.Date(2018, 4, 13, 0, 0, 0, 0, time.UTC)
  fmt.Printf("%v", jdn.ToNumber(t1))
  // Output: 2458222

  t2 := JDN(2448046).ToTime()
  fmt.Printf("%v %v %v\n", t2.Year(), t2.Month(), t2.Day())
  // Output: 1990 June 3
}
```

## More Info

- [Julian day](https://en.wikipedia.org/wiki/Julian_day)
