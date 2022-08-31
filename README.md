<p align="center">
    <a href="https://pkg.go.dev/github.com/kkyr/assert?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="godoc" title="godoc"/></a>
    <a href="https://github.com/kkyr/assert/tags"><img src="https://img.shields.io/github/v/tag/kkyr/assert" alt="semver tag" title="semver tag"/></a>
    <a href="https://goreportcard.com/report/github.com/kkyr/assert"><img src="https://goreportcard.com/badge/github.com/kkyr/assert" alt="go report card" title="go report card"/></a>
    <a href="https://coveralls.io/github/kkyr/assert?branch=main"><img src="https://coveralls.io/repos/github/kkyr/assert/badge.svg?branch=main" alt="coverage status" title="coverage status"/></a>
    <a href="https://github.com/kkyr/assert/blob/main/LICENSE"><img src="https://img.shields.io/github/license/kkyr/assert" alt="license" title="license"/></a>
</p>

# assert

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

<a href="https://pkg.go.dev/github.com/kkyr/assert?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="godoc" title="godoc"/></a>

## Roadmap

- [x] Add assert.Len()
- [ ] Increase UT coverage