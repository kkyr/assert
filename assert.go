package assert

import (
	"bytes"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
	"text/template"
)

// New creates a new Asserter with default options.
func New(t *testing.T) *Asserter {
	asserter := new()
	asserter.t = t
	return asserter
}

func new() *Asserter {
	return &Asserter{
		required: false,
		tmpl:     template.Must(template.New("").Parse(DefaultTmpl)),
		gotText:  DefaultGotText,
		gotVerb:  DefaultGotVerb,
		wantText: DefaultWantText,
		wantVerb: DefaultWantVerb,
		fieldMsg: "",
	}
}

type Asserter struct {
	t        testing.TB
	required bool

	tmpl *template.Template

	gotText  string
	gotVerb  string
	wantText string
	wantVerb string

	fieldMsg string
}

func (a *Asserter) Equal(got any, want any) bool {
	return a.equal(a.t, got, want)
}

func (a *Asserter) Nil(got any) bool {
	return a.nil(a.t, got)
}

func (a *Asserter) Field(field string) *Asserter {
	return a.field(a.t, field)
}

func (a *Asserter) Require() *Asserter {
	return a.require(a.t)
}

func (a *Asserter) copy() *Asserter {
	return &Asserter{
		t:        a.t,
		required: a.required,
		tmpl:     a.tmpl,
		gotText:  a.gotText,
		gotVerb:  a.gotVerb,
		wantText: a.wantText,
		wantVerb: a.wantVerb,
		fieldMsg: a.fieldMsg,
	}
}

func (a *Asserter) equal(t testing.TB, got any, want any) bool {
	t.Helper()
	if !cmp.Equal(got, want) {
		a.fail(t, got, want)
		return false
	}
	return true
}

func (a *Asserter) nil(t testing.TB, got any) bool {
	t.Helper()
	if !isNil(got) {
		a.fail(t, got, nil)
		return false
	}
	return true
}

func (a *Asserter) field(t testing.TB, field string) *Asserter {
	cpy := a.copy()

	cpy.fieldMsg = field
	cpy.t = t

	return cpy
}

func (a *Asserter) require(t testing.TB) *Asserter {
	cpy := a.copy()

	cpy.required = true
	cpy.t = t

	return cpy
}

func (a *Asserter) fail(t testing.TB, got, want any) {
	if a.required {
		t.Fatal(a.format(got, want))
	} else {
		t.Error(a.format(got, want))
	}
}

func (a *Asserter) format(got any, want any) string {
	var buf bytes.Buffer
	a.tmpl.Execute(&buf, a.data())
	return fmt.Sprintf(buf.String(), got, want)
}

func (a *Asserter) data() map[string]string {
	return map[string]string{
		GotTextKey:  a.gotText,
		GotVerbKey:  a.gotVerb,
		WantTextKey: a.wantText,
		WantVerbKey: a.wantVerb,
		FieldValKey: a.fieldMsg,
	}
}

var std = new()

func Equal(t testing.TB, got any, want any) bool {
	return std.equal(t, got, want)
}

func Nil(t testing.TB, got any) bool {
	return std.nil(t, got)
}

func Field(t testing.TB, field string) *Asserter {
	return std.field(t, field)
}

func Require(t testing.TB) *Asserter {
	return std.require(t)
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

// TODO: remove template
// TODO: add assert.Zero()
// TODO: add assert.NotZero()
// TODO: add assert.NotNil()
// TODO: add assert.ErrIs()
// TODO: add assert.Len()
