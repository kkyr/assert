# assert [![Semver Tag](https://img.shields.io/github/v/tag/kkyr/assert)](https://github.com/kkyr/assert/tags) [![Go Report](https://goreportcard.com/badge/github.com/kkyr/assert)](https://goreportcard.com/report/github.com/kkyr/assert) [![Coverage Status](https://coveralls.io/repos/github/kkyr/assert/badge.svg?branch=main)](https://coveralls.io/github/kkyr/assert?branch=main)

assert is a Go library that helps you make assertions in tests, printing a human-readable report of the differences between two values in assertions that fail.

## Installation

```shell
$ go get github.com/kkyr/assert@latest
```

## Usage

```go
package person_test

import (
    "testing"
    
    "github.com/kkyr/assert"
)

func TestPerson(t *testing.T) {
    assert := assert.New(t)

    want := []string{"John", "Jim"}
    got := []string{"John", "Joe"}

    assert.Field("Names").Equal(want, got)
    
    // Output:
    // --- FAIL: TestPerson (0.00s)
    //     person_test.go:15: Names: (-want, +got):
    //              []string{
    //                    "John", 
    //            -       "Jim", 
    //            +       "Joe",
    //              }
}
```

## Documentation

[![GoDoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/kkyr/assert?tab=doc)
