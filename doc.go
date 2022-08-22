// Package assert provides an assertion library for use within Go tests.
//
// # Example Usage
//
//	import (
//		"testing"
//
//		"github.com/kkyr/assert"
//	)
//
//	func TestPeople(t *testing.T) {
//		assert := assert.New(t)
//
//		want, got := "John", "Jane"
//
//		assert.Equal(want, got) // calls t.Error
//	}
//
// If you'd like for assert to call t.Fatal instead of t.Error on failures, use [Assert.Require]:
//
//	func TestPeople(t *testing.T) {
//		require := assert.New(t).Require()
//
//		want, got := "John", "Jane"
//
//		require.Equal(want, got) // calls t.Fatal
//	}
//
// Alternatively, you can have a mixed approach by calling [Assert.Require] only on certain tests using a fluent syntax:
//
//	func TestPeople(t *testing.T) {
//		assert := assert.New(t)
//
//		want, got := "John", "Jane"
//
//		assert.Equal(want, got) // calls t.Error
//		assert.Require().Equal(want, got) // calls t.Fatal
//	}
//
// You can use [Assert.Field] to add additional context to failure messages:
//
//	func TestPeople(t *testing.T) {
//		assert := assert.New(t)
//
//		want, got := "John", "Jane"
//
//		assert.Field("Person").Equal(want, got) // prefixes failure message with "Person: "
//	}
//
// Successive calls to [Assert.Field] will overwrite previous values and only the last field will be used.
//
// You can chain fluent syntax methods together:
//
//	func TestPeople(t *testing.T) {
//		assert := assert.New(t)
//
//		want, got := "John", "Jane"
//
//		assert.Require().Field("Person").Equal(want, got)
//	}
//
// # Equality
//
// Equality of two values is determined using the Equal func in the [go-cmp] package.
//
// [go-cmp]: https://github.com/google/go-cmp
package assert
