package assert

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

// New creates a new Assert that reports failures to tb.
func New(tb testing.TB) *Assert {
	return &Assert{tb: tb}
}

type Assert struct {
	tb      testing.TB
	require bool
	field   string
}

func (a *Assert) Equal(want, got any) bool {
	a.tb.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		msg := fmt.Sprintf("(-want, +got):\n%s", diff)
		a.fail(a.addField(msg))
		return false
	}

	return true
}

func (a *Assert) NotEqual(want, got any) bool {
	a.tb.Helper()

	if cmp.Equal(want, got) {
		a.fail(a.format("equal values", "non-equal values"))
		return false
	}

	return true
}

func (a *Assert) Nil(value any) bool {
	a.tb.Helper()

	if !isNil(value) {
		a.fail(a.format(nil, value))
		return false
	}

	return true
}

func (a *Assert) NotNil(value any) bool {
	a.tb.Helper()

	if isNil(value) {
		a.fail(a.format("<non-nil>", "<nil>"))
		return false
	}

	return true
}

func (a *Assert) ErrorIs(err, target error) bool {
	a.tb.Helper()

	if !errors.Is(err, target) {
		a.fail(a.addField("no error in err's chain matches target"))
		return false
	}

	return true
}

func (a *Assert) Zero(value any) bool {
	a.tb.Helper()

	if !isZero(value) {
		a.fail(a.format("zero value", value))
		return false
	}

	return true
}

func (a *Assert) NotZero(value any) bool {
	a.tb.Helper()

	if isZero(value) {
		a.fail(a.format("non-zero value", value))
		return false
	}

	return true
}

func (a *Assert) Field(value string) *Assert {
	assertCpy := a.copy()
	assertCpy.field = value
	return assertCpy
}

func (a *Assert) Require() *Assert {
	assertCpy := a.copy()
	assertCpy.require = true
	return assertCpy
}

func (a *Assert) copy() *Assert {
	return &Assert{
		tb:      a.tb,
		require: a.require,
		field:   a.field,
	}
}

func (a *Assert) fail(args ...any) {
	a.tb.Helper()
	if a.require {
		a.tb.Fatal(args...)
	} else {
		a.tb.Error(args...)
	}
}

func (a *Assert) format(want, got any) string {
	return a.addField(fmt.Sprintf("want %v, got %v", want, got))
}

func (a *Assert) addField(s string) string {
	if a.field != "" {
		s = fmt.Sprintf("%s: %s", a.field, s)
	}
	return s
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

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()

	nilable := []reflect.Kind{reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice}
	for _, k := range nilable {
		if kind == k {
			return value.IsNil()
		}
	}

	return false
}

// TODO: add assert.Len()
