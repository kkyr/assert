package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// New returns a new instance of Assert that reports any assertion
// failures to tb.
func New(tb testing.TB) *Assert {
	return &Assert{tb: tb}
}

// Assert is a type that can make assertions.
//
// By default, assertion failures are reported using t.Error.
type Assert struct {
	tb testing.TB

	// if set to true will call t.Fatal instead of t.Error on failures.
	require bool

	// if set will prefix this in failure messages.
	field string
}

// Equal asserts that want and got are equal.
// See [cmp.Equal] for details on how equality is determined.
func (a *Assert) Equal(want, got any) bool {
	a.tb.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		a.fail(fmt.Sprintf("(-want, +got):\n%s", diff))
		return false
	}

	return true
}

// NotEqual asserts that want and got are not equal.
// See [cmp.Equal] for details on how equality is determined.
func (a *Assert) NotEqual(want, got any) bool {
	a.tb.Helper()

	if cmp.Equal(want, got) {
		a.fail(a.format("equal values", "non-equal values"))
		return false
	}

	return true
}

// Nil asserts that value is nil.
func (a *Assert) Nil(value any) bool {
	a.tb.Helper()

	if !isNil(value) {
		a.fail(a.format(nil, value))
		return false
	}

	return true
}

// NotNil asserts that value is not nil.
func (a *Assert) NotNil(value any) bool {
	a.tb.Helper()

	if isNil(value) {
		a.fail(a.format("<non-nil>", "<nil>"))
		return false
	}

	return true
}

// ErrorIs asserts that at least one of the error in err's chain matches target.
// See [errors.Is] for details on how a matching error is found.
func (a *Assert) ErrorIs(err, target error) bool {
	a.tb.Helper()

	if !errors.Is(err, target) {
		a.fail("no error in err's chain matches target")
		return false
	}

	return true
}

// Zero asserts that value is the zero value for its type.
// Pointer values are determined based on the zero value of the referenced values.
func (a *Assert) Zero(value any) bool {
	a.tb.Helper()

	if !isZero(value) {
		a.fail(a.format("zero value", value))
		return false
	}

	return true
}

// NotZero asserts that value is not the zero value for its type.
// Pointer values are determined based on the zero value of the referenced values.
func (a *Assert) NotZero(value any) bool {
	a.tb.Helper()

	if isZero(value) {
		a.fail(a.format("non-zero value", value))
		return false
	}

	return true
}

// Len asserts that value has length n.
// If len() cannot be applied to value, the test fails.
func (a *Assert) Len(value any, n int) bool {
	a.tb.Helper()

	got, ok := getLen(value)
	if !ok {
		a.fail(fmt.Sprintf("could not apply len() to %T", value))
		return false
	}

	if got != n {
		a.fail(fmt.Sprintf("want len() = %v, got %v", n, got))
		return false
	}

	return true
}

// Field returns a copy of Assert that will prefix failure messages with s.
//
//	assert.Field("Age").Equal(18, 20)
//
// This should be used to enrich failure messages with information about the
// field that is being asserted.
func (a *Assert) Field(s string) *Assert {
	cpy := a.copy()
	cpy.field = s
	return cpy
}

// Require returns a copy of Assert that will call t.Fatal on failures.
func (a *Assert) Require() *Assert {
	cpy := a.copy()
	cpy.require = true
	return cpy
}

func (a *Assert) copy() *Assert {
	return &Assert{
		tb:      a.tb,
		require: a.require,
		field:   a.field,
	}
}

func (a *Assert) fail(msg string) {
	a.tb.Helper()

	if a.field != "" {
		msg = fmt.Sprintf("%s: %s", a.field, msg)
	}

	if a.require {
		a.tb.Fatal(msg)
	} else {
		a.tb.Error(msg)
	}
}

func (a *Assert) format(want, got any) string {
	return fmt.Sprintf("want %v, got %v", want, got)
}

func isZero(value any) bool {
	if value == nil {
		return true
	}

	if i, ok := value.(interface{ IsZero() bool }); ok {
		return i.IsZero()
	}

	switch rv := reflect.ValueOf(value); rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice:
		return rv.Len() == 0
	case reflect.Ptr:
		if rv.IsNil() {
			return true
		}
		// dereference pointer and recursively test it
		return isZero(rv.Elem().Interface())
	default:
		return rv.IsZero()
	}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	kind := rv.Kind()

	nilable := []reflect.Kind{reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice}
	for _, k := range nilable {
		if kind == k {
			return rv.IsNil()
		}
	}

	return false
}

func getLen(value any) (int, bool) {
	rv := reflect.ValueOf(value)

	switch k := rv.Kind(); k {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len(), true
	case reflect.Ptr:
		if rv.Type().Elem().Kind() == reflect.Array {
			return rv.Type().Elem().Len(), true
		}
	}

	return 0, false
}
