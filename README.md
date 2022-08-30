# assert

assert is a Go library for making assertions in tests.

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

- [ ] Add assert.Len()
- [ ] Increase UT coverage