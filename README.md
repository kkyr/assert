## Installation

Add the package to your module:

```shell
$ go get -u github.com/kkyr/assert@latest
```

## Usage

### Basic

Import the package into your tests and start using it:

```go
import (
	"testing"
    "github.com/kkyr/assert"
)

func TestSuperhero(t *testing.T) {
	got, want := "batman", "batwoman"
	assert.Equal(t, got, want)
	// Output: 
	// got batman, want batwoman
}
```

### Field name

You can optionally specify a field name for additional context during failures:
	
```go
func TestSuperhero(t *testing.T) {
	got, want := "batman", "batwoman"
	assert.Field("Superhero").Equal(t, got, want)
	// Output: 
	// Superhero: got batman, want batwoman
}
```

### Failure mode

By default validation failures will call t.Fatal() which will end the current test. If instead you'd like tests to continue execution after a failure, configure the assertion with NoFatal():

```go
assert.NoFatal().Equal(got, want)
```

Which will report failures with t.Error() instead of t.Fatal().

## Customization

Do you prefer the terminology "actual" and "expected" over "got" and "want"? No problem:

```go
func TestSuperhero(t *testing.T) {
    assert.SetGotText("actual")
    assert.SetWantText("expected")
    
    got, want := "batman", "batwoman"
    assert.Equal(t, got, want)
    // Output: 
    // actual batman, expected batwoman
}
```

Would you like failures to emit quoted strings instead? Say no more:

```go
func TestSuperhero(t *testing.T) {
    assert.SetGotVerb("%q")
    assert.SetWantVerb("%q")
    
    got, want := "batman", "batwoman"
    assert.Equal(t, got, want)
    // Output: 
    // got "batman", want "batwoman"
}
```

Or just don't like the template? Define your own:

```go
func TestSuperhero(t *testing.T) {
    template := fmt.Sprintf("{{.%s}}={{.%s}}, {{.%s}} {{.%s}}",
    assert.FieldValKey, assert.GotVerbKey, assert.WantTextKey, assert.WantVerbKey)
    assert.SetTemplate(template) // returns err if template parsing fails
    
    got, want := "batman", "batwoman"
    assert.Field("Superhero").Equal(t, got, want)
	// Superhero=batman, want batwoman
}
```